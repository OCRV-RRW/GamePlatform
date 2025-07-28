package user

import (
	"gameplatform/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, handler *UserHandler, userMiddleware *middleware.UserMiddleware) {
	user := router.Group("/users")
	me := user.Group("/me")

	me.Get("/", userMiddleware.DeserializeUser, handler.GetMe)
	me.Patch("/", userMiddleware.DeserializeUser, handler.UpdateMe)

	user.Get("/", userMiddleware.DeserializeUser, handler.GetUser)
	user.Patch("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, middleware.AdminUser, handler.UpdateUser)
	user.Delete("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, middleware.AdminUser, handler.DeleteUser)
}
