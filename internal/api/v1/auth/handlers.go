package auth

import (
	"context"
	"errors"
	"fmt"
	"gameplatform/internal/DTO"
	"gameplatform/internal/api"
	"gameplatform/internal/config"
	"gameplatform/internal/dbconn"
	"gameplatform/internal/repository"
	"gameplatform/internal/token"
	"gameplatform/internal/utils"
	"gameplatform/internal/validation"
	"strings"
	"time"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type AuthHandler struct {
	Config     *config.Config
	SMTP       *utils.SMTP
	Repository *repository.UserRepository
	Redis      *dbconn.RedisConnection
	conv       DTO.UserConverter
}

func NewAuthHandler(repository *repository.UserRepository, redis *dbconn.RedisConnection, smtp *utils.SMTP, config *config.Config) AuthHandler {
	return AuthHandler{
		Repository: repository,
		Redis:      redis,
		Config:     config,
		SMTP:       smtp,
		conv:       &DTO.UserConverterImpl{},
	}
}

// SignUpUser godoc
//
// @Description	 sign up user
// @Tags       Auth
// @Accept		 json
// @Produce		 json
// @Param        SignUpInput		body		DTO.SignUpInput		true   "SignUpInput"
// @Success		 201 {object} api.SuccessResponse[DTO.UserResponseDTO]
// @Failure      400 {object} api.ErrorResponse
// @Failure      409 {object} api.ErrorResponse
// @Failure      500 {object} api.ErrorResponse
// @Failure      422 {object} api.ErrorResponse
// @Router		 /api/v1/auth/register [post]
func (h *AuthHandler) SignUpUser(c *fiber.Ctx) error {
	var payload *DTO.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.UnprocessableEntity, Message: err.Error()},
		}))
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse(userErrors))
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.IncorrectParameter, Parameter: "password", Message: "password and confirm password don't match"},
			{Code: api.IncorrectParameter, Parameter: "confirm_password", Message: "password and confirm password don't match"},
		}))
	}

	hashedPassword, err := utils.HashPassword(payload.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	verificationCode := utils.GenerateCode(20)

	newUser := repository.CreateUser{
		Name:             payload.Name,
		Email:            strings.ToLower(payload.Email),
		Password:         hashedPassword,
		Verified:         false,
		VerificationCode: verificationCode,
	}

	sendEmailError := errors.New("Something went wrong on sending email")

	user, err := h.Repository.CreateUserWithTransaction(newUser, func(user repository.CreateUser) error {
		//Send verification code.
		err = h.SMTP.SendEmail(newUser.Email, &utils.EmailData{
			URL:       h.Config.ClientOrigin + "/register/verify/" + verificationCode,
			FirstName: newUser.Name,
			Subject:   "Your account verification code",
		}, "verificationCode.html")
		if err != nil {
			return sendEmailError
		}
		return nil
	})

	if errors.Is(err, sendEmailError) {
		param := validation.GetJSONTag(payload, "Email")
		userErrors = []*api.Error{validation.GetErrorResponse(param, "email")}
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse(userErrors))
	} else if errors.Is(err, repository.ErrDuplicatedKey) {
		return c.Status(fiber.StatusConflict).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.EmailAlreadyExists, Message: "email already exists"},
		}))
	} else if err != nil {
		return api.InternalServerError(c, err, "Something went wrong on sending email")
	}

	return c.Status(fiber.StatusCreated).JSON(api.NewSuccessResponse(
		DTO.UserResponseDTO{User: h.conv.GetUserToUserResponse(user)},
		"We sent an email with a verification code to "+newUser.Email))
}

// SignInUser godoc
//
// @Description	sign in user
// @Tags        Auth
// @Accept      json
// @Param       SignInInput		body		DTO.SignInInput		true   "SignInInput"
// @Produce		json
// @Success		200 {object} api.SuccessResponse[DTO.TokenResponse]
// @Failure     400 {object} api.ErrorResponse
// @Failure     422 {object} api.ErrorResponse
// @Router	    /api/v1/auth/login [post]
func (h *AuthHandler) SignInUser(c *fiber.Ctx) error {
	var payload *DTO.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return api.BadRequestParamError(c, userErrors)
	}

	user, err := h.Repository.GetByEmail(payload.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.InvalidEmailOrPassword, Message: "invalid email or password"},
		}))
	}

	err = utils.VerifyPassword([]byte(user.Password), payload.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.InvalidEmailOrPassword, Message: "invalid email or password"},
		}))
	}

	if !user.Verified {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.InvalidEmailOrPassword, Message: "invalid email or password"},
		}))
	}

	return h.generateAndSendToken(c, user, "success authorize")
}

// LogoutUser godoc
//
// @Description	logout
// @Tags         Auth
// @Accept		json
// @Produce		json
// @Success	    200 {object} api.Response
// @Failure     403 {object} api.ErrorResponse
// @Failure     500 {object} api.ErrorResponse
// @Router		/api/v1/auth/logout [get]
func (h *AuthHandler) LogoutUser(c *fiber.Ctx) error {
	message := "Token is invalid or session has expired"

	refresh_token := c.Cookies("refresh_token")
	if refresh_token == "" {
		return api.ForbiddenError(c, message)
	}

	tokenClaims, err := token.ValidateToken(refresh_token, h.Config.RefreshTokenPublicKey)
	if err != nil {
		return api.ForbiddenError(c, message)
	}

	accessTokenUuid := c.Locals("access_token_uuid").(string)
	tokenRepo := token.NewAuthTokenRepository(h.Redis.RedisClient)
	err = tokenRepo.RemoveTokenByTokenUuid(accessTokenUuid, tokenClaims.TokenUuid)
	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Domain:   h.Config.CookieDomain,
		Expires:  expired,
		SameSite: "none",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Domain:   h.Config.CookieDomain,
		Expires:  expired,
		SameSite: "none",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Value:    "",
		Domain:   h.Config.CookieDomain,
		Expires:  expired,
		SameSite: "none",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// RefreshAccessToken godoc
//
// @Description	refresh access token
// @Tags        Auth
// @Accept		json
// @Success	    200 {object} api.SuccessResponse[DTO.TokenResponse]
// @Failure     403 {object} api.ErrorResponse
// @Failure     500 {object} api.ErrorResponse
// @Failure     422 {object} api.ErrorResponse
// @Router		/api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshAccessToken(c *fiber.Ctx) error {
	message := "could not refresh access token"

	refresh_token := c.Cookies("refresh_token")

	if refresh_token == "" {
		return api.ForbiddenError(c, message)
	}

	// Validate refresh token, and search token in redis
	tokenClaims, err := token.ValidateToken(refresh_token, h.Config.RefreshTokenPublicKey)
	if err != nil {
		return api.ForbiddenError(c, message)
	}
	tokenRepo := token.NewAuthTokenRepository(h.Redis.RedisClient)
	userid, err := tokenRepo.GetUserIdByTokenUuid(tokenClaims.TokenUuid)
	if errors.Is(err, redis.Nil) || userid == "" {
		return api.ForbiddenError(c, message)
	}

	user, err := h.Repository.GetUserById(userid)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.ForbiddenError(c, "the user belonging to this token no logger exists")
		} else {
			return api.InternalServerError(c, fmt.Errorf("Couldn't get user by token: %v", err.Error()), "something went wrong")
		}
	}

	// Remove old token
	err = tokenRepo.RemoveTokenByTokenUuid(tokenClaims.TokenUuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.TokenInvalidOrExpired, Message: "invalid token"},
		}))
	}

	return h.generateAndSendToken(c, user, "success refresh token")
}

// VerifyEmail godoc
//
// @Description	 verify user email
// @Tags         Auth
// @Produce      json
// @Param        verify_code   path string true "Verification code"
// @Success		 200 {object} api.Response
// @Failure      400 {object} api.ErrorResponse
// @Failure      409 {object} api.ErrorResponse
// @Router		 /api/v1/auth/verify-email [post]
func (h *AuthHandler) VerifyEmail(c *fiber.Ctx) error {
	verificationCode := c.Params("verificationCode")

	user, err := h.Repository.GetByVerificationCode(verificationCode)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.InvalidVerificationCode, Message: "Invalid verification code or user doesn't exists"},
		}))
	}

	if user.Verified {
		return c.Status(fiber.StatusConflict).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.InvalidVerificationCode, Message: "email already verified"},
		}))
	}

	err = h.Repository.UpdateUserVerification(user.ID, "", true)
	if err != nil {
		slog.ErrorContext(c.Context(), err.Error())
		return api.InternalServerError(c, err, "something went wrong")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Email verified successfully"})
}

// ForgotPassword godoc
//
// @Description	 forgot password
// @Tags         Auth
// @Accept		 json
// @Produce      json
// @Param        ForgotPasswordInput		body		DTO.ForgotPasswordInput		true   "ForgotPasswordInput"
// @Success		 200 {object} api.Response
// @Failure      400 {object} api.ErrorResponse
// @Failure      401 {object} api.ErrorResponse  "User email is not verified"
// @Failure      404 {object} api.ErrorResponse
// @Router		 /api/v1/auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	var payload DTO.ForgotPasswordInput

	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return api.BadRequestParamError(c, userErrors)
	}

	user, err := h.Repository.GetByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.IncorrectParameter, Parameter: "email", Message: "invalid email"},
		}))
	}

	if !user.Verified {
		return api.UnauthorizedError(c, "unauthorized")
	}

	// Generate and send to email verification Code
	resetToken := utils.GenerateCode(30)
	ctx := context.TODO()
	h.Redis.RedisClient.Set(ctx, resetToken, user.ID.String(), h.Config.ResetPasswordTokenExpiredIn)
	h.SMTP.SendEmail(user.Email, &utils.EmailData{
		URL:       h.Config.ClientOrigin + "/forgot_password/reset/" + resetToken,
		FirstName: user.Name,
		Subject:   "Your password reset token (valid for 10min)",
	}, "resetPassword.html")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "We sent an email with a reset code to " + user.Email,
	})
}

// ResetPassword godoc
//
// @Description	 reset user password
// @Tags         Auth
// @Accept		 json
// @Produce		 json
// @Param        reset_code   path string true "reset code"
// @Param        ResetPasswordInput		body		DTO.ResetPasswordInput		true   "ResetPasswordInput"
// @Success		 200 {object} api.Response
// @Failure      400 {object} api.ErrorResponse
// @Failure      500 {object} api.ErrorResponse
// @Router		 /api/v1/auth/reset-password [patch]
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	var payload *DTO.ResetPasswordInput
	resetToken := c.Params("resetToken")

	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}
	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return api.BadRequestParamError(c, userErrors)
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.IncorrectParameter, Parameter: "password", Message: "password and password_confirm do not match"},
			{Code: api.IncorrectParameter, Parameter: "password_confirm", Message: "password and password_confirm do not match"},
		}))
	}

	tokenRepo := token.NewAuthTokenRepository(h.Redis.RedisClient)

	ctx := context.Background()
	userid, err := h.Redis.RedisClient.Get(ctx, resetToken).Result()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.TokenInvalidOrExpired, Message: "The reset token is invalid or has expired"},
		}))
	}
	user, err := h.Repository.GetUserById(userid)
	if err != nil {
		return api.NotFoundError(c, "couldn't find the user")
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	err = h.Repository.UpdatePassword(user.ID, string(hashedPassword))

	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	_, err = h.Redis.RedisClient.Del(ctx, resetToken).Result()
	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	err = tokenRepo.RemoveAllUserToken(user.ID.String())
	if err != nil {
		return api.InternalServerError(c, fmt.Errorf("Couldn't reset user token error: %v", err), "something went wrong")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password data updated successfully"})
}

func (h *AuthHandler) generateAndSendToken(c *fiber.Ctx, user *repository.GetUser, message string) error {
	tokenRepo := token.NewAuthTokenRepository(h.Redis.RedisClient)

	// Create and save new access token
	now := time.Now()
	accessTokenDetails, err := token.CreateToken(user.ID.String(), h.Config.AccessTokenExpiresIn, h.Config.AccessTokenPrivateKey)
	if err != nil {
		return api.InternalServerError(c, fmt.Errorf("Something went wrong when creating the access token %v", err), "something went wrong")
	}

	refreshTokenDetails, err := token.CreateToken(user.ID.String(), h.Config.RefreshTokenExpiresIn, h.Config.RefreshTokenPrivateKey)
	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	errRefresh := tokenRepo.SaveToken(
		user.ID.String(),
		refreshTokenDetails,
		time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now))

	if errRefresh != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Domain:   h.Config.CookieDomain,
		Path:     "/",
		MaxAge:   h.Config.AccessTokenMaxAge * 60,
		Secure:   h.Config.CookieSecure,
		HTTPOnly: true,
		SameSite: "none",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "logged_in",
		Domain:   h.Config.CookieDomain,
		Value:    "true",
		Path:     "/",
		MaxAge:   h.Config.AccessTokenMaxAge * 60,
		Secure:   h.Config.CookieSecure,
		HTTPOnly: false,
		SameSite: "none",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Domain:   h.Config.CookieDomain,
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   h.Config.RefreshTokenMaxAge * 60,
		Secure:   h.Config.CookieSecure,
		HTTPOnly: true,
		SameSite: "none",
	})

	expiredIn := time.Unix(*accessTokenDetails.ExpiresIn, 0)
	return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(
		DTO.TokenResponse{
			AccessToken: *accessTokenDetails.Token,
			ExpiredIn:   &expiredIn,
		},
		message))
}
