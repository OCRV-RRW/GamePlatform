package auth

import (
	"gameplatform/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, handler *AuthHandler, userMiddleware *middleware.UserMiddleware) {
	auth := router.Group("auth")
	auth.Post("/register", handler.SignUpUser)
	auth.Post("/refresh", handler.RefreshAccessToken)
	auth.Post("/login", handler.SignInUser)
	auth.Get("/logout", userMiddleware.DeserializeUser, handler.LogoutUser)
	auth.Post("/verify-email/:verificationCode", handler.VerifyEmail)
	auth.Post("/forgot-password", handler.ForgotPassword)
	auth.Patch("/reset-password/:resetToken", handler.ResetPassword)
}
