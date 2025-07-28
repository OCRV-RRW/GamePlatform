package view

import "github.com/gofiber/fiber/v2"

func AddRoutes(router fiber.Router, viewHandler *ViewHandler) {
	view := router.Group("/")

	view.Get("/", viewHandler.Home)
	view.Get("/game/:id", viewHandler.Show)
	view.Get("/play/:id", viewHandler.Game)
}
