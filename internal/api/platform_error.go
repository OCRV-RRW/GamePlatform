package api

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

var (
	IncorrectParameter      = "INCORRECT_PARAMETER"
	ServerError             = "SERVER_ERROR"
	NotFound                = "NOT_FOUND"
	InvalidEmailOrPassword  = "INVALID_EMAIL_OR_PASSWORD"
	UnprocessableEntity     = "UNPROCESSABLE_ENTITY"
	EmailAlreadyExists      = "EMAIL_ALREADY_EXISTS"
	InvalidVerificationCode = "INVALID_VERIFICATION_CODE"
	Forbidden               = "FORBIDDEN"
	Unauthorized            = "UNAUTHORIZED"
	TokenInvalidOrExpired   = "TOKEN_INVALID_OR_EXPIRED"
)

type SuccessResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Error struct {
	Code      string `json:"code"`
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}

type ErrorResponse struct {
	Errors []*Error `json:"error"`
	Status string   `json:"status"`
}

func NewErrorResponse(errors []*Error) *ErrorResponse {
	return &ErrorResponse{
		Errors: errors,
		Status: "fail",
	}
}

func NewSuccessResponse[T any](data T, message string) *SuccessResponse[T] {
	return &SuccessResponse[T]{
		Data:    data,
		Status:  "success",
		Message: message,
	}
}

// Error Halpers

func NotFoundError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse([]*Error{
		{Code: NotFound, Message: message},
	}))
}

func InternalServerError(c *fiber.Ctx, err error, userMessage string) error {
	slog.ErrorContext(c.Context(), err.Error())
	return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse([]*Error{
		{Code: ServerError, Message: userMessage},
	}))
}

func UnprocessableEntityError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(NewErrorResponse([]*Error{
		{Code: UnprocessableEntity, Message: err.Error()},
	}))
}

func BadRequestParamError(c *fiber.Ctx, errors []*Error) error {
	return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse(errors))
}

func ForbiddenError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(NewErrorResponse([]*Error{
		{Code: Forbidden, Message: message},
	}))
}

func UnauthorizedError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(NewErrorResponse([]*Error{
		{Code: Unauthorized, Message: message},
	}))
}
