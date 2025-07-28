package view

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func InternalServerErrorDefault(c *fiber.Ctx, err error) error {
	slog.ErrorContext(c.Context(), err.Error())
	_, err = c.Status(fiber.StatusInternalServerError).WriteString("something went wrong")
	return err
}

func InternalServerError(c *fiber.Ctx, err error, userMessage string) error {
	slog.ErrorContext(c.Context(), err.Error())
	_, err = c.Status(fiber.StatusInternalServerError).WriteString(userMessage)
	return err
}

func NotFoundError(c *fiber.Ctx, userMessage string) error {
	_, err := c.Status(fiber.StatusNotFound).WriteString(userMessage)
	return err
}
