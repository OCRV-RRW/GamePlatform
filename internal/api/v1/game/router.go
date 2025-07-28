package game

import (
	"gameplatform/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(router fiber.Router, handler *GameHandler, userMiddleware *middleware.UserMiddleware) {
	game := router.Group("/games")
	preview := game.Group("/previews")

	preview.Delete("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, handler.DeletePreview)
	preview.Post("/", userMiddleware.DeserializeUser, middleware.AdminUser, handler.CreatePreview)

	game.Get("/", userMiddleware.DeserializeUser, handler.GetGames)
	game.Get("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, handler.GetGame)
	game.Post("/", userMiddleware.DeserializeUser, middleware.AdminUser, handler.CreateGame)
	game.Patch("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, middleware.AdminUser, handler.UpdateGame)
	game.Delete("/:id", middleware.CheckUUID, userMiddleware.DeserializeUser, middleware.AdminUser, handler.DeleteGame)
}
