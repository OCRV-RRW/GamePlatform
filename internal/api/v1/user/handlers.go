package user

import (
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
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Config     *config.Config
	SMTP       *utils.SMTP
	Repository *repository.UserRepository
	Redis      *dbconn.RedisConnection
	Converter  DTO.UserConverter
}

func NewUserHandler(
	conf *config.Config,
	smtp *utils.SMTP,
	repository *repository.UserRepository,
	redis *dbconn.RedisConnection,
	converter DTO.UserConverter) *UserHandler {
	return &UserHandler{
		Config:     conf,
		SMTP:       smtp,
		Repository: repository,
		Redis:      redis,
		Converter:  converter,
	}
}

// GetMe godoc
//
// @Description	 get current user
// @Tags         User
// @Produce		 json
// @Success		 200 {object}  api.SuccessResponse[DTO.UserResponseDTO]
// @Failure      401
// @Router		 /api/v1/users/me [get]
func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*repository.GetUser)

	if !ok {
		return api.InternalServerError(c, fmt.Errorf("coudn't get the user form locals"), "somthing went wrong")
	}

	return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(
		DTO.UserResponseDTO{User: h.Converter.GetUserToUserResponse(user)}, ""))
}

// DeleteUser godoc
//
// @Description	 delete user by id
// @Tags         User
// @Produce		 json
// @Param        id   path string true "User ID"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/users/ [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := h.Repository.DeleteUser(userId)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found user")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	tokenRepo := token.NewAuthTokenRepository(h.Redis.RedisClient)
	err = tokenRepo.RemoveAllUserToken(userId)
	if err != nil {
		slog.ErrorContext(c.Context(), fmt.Sprintf("Couldn't reset user token error: %v", err))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// GetUser godoc
//
// @Description	 get user by id
// @Tags         User
// @Produce		 json
// @Param        id  query     string     false  "user id"
// @Success		 200 {object} api.SuccessResponse[DTO.UsersResponse]
// @Failure      500
// @Router		 /api/v1/users/ [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Query("id")
	if id != "" {
		userResponse, err := h.getUserById(c, id)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(fiber.Map{
			"users": []*DTO.UserResponse{userResponse}}, ""))
	}

	email := c.Query("email")
	if email != "" {
		userResponse, err := h.getUserByEmail(c, email)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(fiber.Map{
			"users": []*DTO.UserResponse{userResponse}}, ""))
	}

	users, err := h.Repository.GetAll()
	if err != nil {
		return api.InternalServerError(c, err, "couldn't get users")
	}

	userRecords := h.Converter.GetUsersToUserResponses(users)

	return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(fiber.Map{"users": userRecords}, ""))
}

// UpdateMe godoc
//
// @Description	 update user
// @Tags         User
// @Produce		 json
// @Param        UpdateUserInput		body		DTO.UpdateUserInput		true   "UpdateUserInput"
// @Success		 200
// @Failure      500
// @Failure      404
// @Router		 /api/v1/users/me [patch]
func (h *UserHandler) UpdateMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*repository.GetUser)
	return h.updateUserByPayload(c, user)
}

// UpdateUser godoc
//
// @Description	 update another user
// @Tags         User
// @Produce		 json
// @Param        id   path string true "User ID"
// @Param        UpdateUserInput		body		DTO.UpdateUserInput		true   "UpdateUserInput"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/users/ [patch]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	user, err := h.Repository.GetUserById(c.Params("id"))

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found user")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	return h.updateUserByPayload(c, user)
}

func (h *UserHandler) updateUserByPayload(c *fiber.Ctx, old_user *repository.GetUser) error {
	var payload *DTO.UpdateUserInput
	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}

	userErrors := validation.ValidateStruct(payload)
	if userErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse(userErrors))
	}

	update_user := h.Converter.GetUserToUpdateUser(old_user)

	if payload.Birthday != nil {
		if payload.Birthday.Day() > 1 {
			update_user.Birthday = payload.Birthday
		}
	}
	if payload.Gender != nil && (*payload.Gender == "М" || *payload.Gender == "Ж") {
		update_user.Gender = payload.Gender
	}
	if payload.Name != nil {
		update_user.Name = *payload.Name
	}
	if payload.IsAdmin != nil {
		if !old_user.IsAdmin && *payload.IsAdmin {
			return api.ForbiddenError(c, "Permission denied")
		}

		update_user.IsAdmin = *payload.IsAdmin
	}
	err := h.Repository.Update(&update_user)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found user")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func (h *UserHandler) getUserByEmail(c *fiber.Ctx, email string) (*DTO.UserResponse, error) {
	user, err := h.Repository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, api.NotFoundError(c, "not found user")
		} else {
			return nil, api.InternalServerError(c, err, "something went wrong")
		}
	}

	userResponse := h.Converter.GetUserToUserResponse(user)
	return &userResponse, nil
}

func (h *UserHandler) getUserById(c *fiber.Ctx, id string) (*DTO.UserResponse, error) {
	user, err := h.Repository.GetUserById(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, api.NotFoundError(c, "not found user")
		} else {
			return nil, api.InternalServerError(c, err, "something went wrong")
		}
	}

	userResponse := h.Converter.GetUserToUserResponse(user)
	return &userResponse, nil
}
