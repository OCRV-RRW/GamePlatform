package middleware

import (
	"errors"
	"fmt"
	"gameplatform/internal/api"
	"gameplatform/internal/config"
	"gameplatform/internal/repository"
	"gameplatform/internal/token"
	"gameplatform/internal/utils"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type UserMiddleware struct {
	Repository *repository.UserRepository
	Config     *config.Config
}

func NewUserMiddleware(repository *repository.UserRepository, config *config.Config) UserMiddleware {
	return UserMiddleware{
		Repository: repository,
		Config:     config,
	}
}

func (m *UserMiddleware) DeserializeUser(c *fiber.Ctx) error {
	access_token := utils.GetToken(c)

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenClaims, err := token.ValidateToken(access_token, m.Config.AccessTokenPublicKey)

	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// Check token in redis
	//tokenRepo := token.NewAuthTokenRepository(database.RedisClient)
	//userid, err := tokenRepo.GetUserIdByTokenUuid(tokenClaims.TokenUuid)
	//if err != nil {
	//	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "Token is invalid or session has expired"})
	//}

	user, err := m.Repository.GetUserById(tokenClaims.UserID)

	if errors.Is(err, repository.ErrRecordNotFound) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	} else if err != nil {
		slog.ErrorContext(c.Context(), fmt.Sprintf("Couldn't get user by id in middelware. Error: %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "something went wrong"})
	}

	c.Locals("user", user)
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return c.Next()
}

func AdminUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*repository.GetUser)
	if user.IsAdmin {
		return c.Next()
	}

	return c.Status(fiber.StatusForbidden).JSON(*api.NewErrorResponse([]*api.Error{
		{Code: api.Forbidden, Message: "Permission denied"},
	}))
}
