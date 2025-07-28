package view

import (
	"errors"
	"gameplatform/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type ViewHandler struct {
	Repository *repository.GameRepository
	conv       ViewConverter
}

func NewViewHanlder(repository *repository.GameRepository, conv ViewConverter) ViewHandler {
	return ViewHandler{
		Repository: repository,
		conv:       conv,
	}
}

func (h *ViewHandler) Home(c *fiber.Ctx) error {
	games, err := h.Repository.GetGameAll()
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return NotFoundError(c, "not found games")
		} else {
			return InternalServerErrorDefault(c, err)
		}

	}
	result := h.conv.GetGamesToGameViews(games)
	data := HomePageData{
		Games: result,
	}

	return c.Render("home.page.html", data)
}

func (h *ViewHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	game, err := h.Repository.GetGameByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return NotFoundError(c, "not found game")
		} else {
			return InternalServerErrorDefault(c, err)
		}
	}

	previews, err := h.Repository.GetPreviews(id)
	if err != nil && !errors.Is(err, repository.ErrRecordNotFound) {
		return InternalServerErrorDefault(c, err)
	}

	gameResult := h.conv.GetGameToGameView(game)
	previewsResult := h.conv.GetPreviewsToPreviewViews(previews)
	result := GamePreivewView{
		GameView: gameResult,
		Preview:  previewsResult,
	}

	data := ShowPageData{
		Game: result,
	}

	return c.Render("show.page.html", data, "base")
}

func (h *ViewHandler) Game(c *fiber.Ctx) error {
	id := c.Params("id")

	game, err := h.Repository.GetGameByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return NotFoundError(c, "not found game")
		} else {
			return InternalServerErrorDefault(c, err)
		}
	}

	gameResult := h.conv.GetGameToGameView(game)

	data := GamePageData{
		Game: gameResult,
	}

	return c.Render("game.page.html", data, "base")
}
