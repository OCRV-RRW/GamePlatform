package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CheckUUID(c *fiber.Ctx) error {
	id := c.Params("id")

	if strings.EqualFold(id, "") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "id is empty"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "incorrect uuid"})
	}

	return c.Next()
}
