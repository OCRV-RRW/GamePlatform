package game

import (
	"errors"
	"fmt"
	"gameplatform/internal/DTO"
	"gameplatform/internal/api"
	"gameplatform/internal/config"
	"gameplatform/internal/dbconn"
	"gameplatform/internal/repository"
	"gameplatform/internal/utils"
	"gameplatform/internal/validation"
	"log/slog"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type GameHandler struct {
	Config     *config.Config
	Repository *repository.GameRepository
	Redis      *dbconn.RedisConnection
	Minio      *dbconn.MinioConnection
	Converter  DTO.GameConverter
}

func NewGameHandler(
	conf *config.Config,
	repository *repository.GameRepository,
	redis *dbconn.RedisConnection,
	minio *dbconn.MinioConnection,
	converter DTO.GameConverter) *GameHandler {
	return &GameHandler{
		Config:     conf,
		Repository: repository,
		Minio:      minio,
		Redis:      redis,
		Converter:  converter,
	}
}

// GetGame godoc
//
// @Description  get game by id
// @Tags         Game
// @Produce      json
// @Param        id   path string true "Game ID"
// @Success      200 {object} api.SuccessResponse[DTO.GameResponseDTO]
// @Failure      500
// @Router       /api/v1/games/{id} [get]
func (h *GameHandler) GetGame(c *fiber.Ctx) error {
	id := c.Params("id")

	game, err := h.Repository.GetGameByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found game")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	previews, err := h.Repository.GetPreviews(id)
	if err != nil {
		if !errors.Is(err, repository.ErrRecordNotFound) {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}
	preview_responses := h.Converter.GetPreviewsToPreviewResponses(previews)

	game_response := h.Converter.GetGameToGameResponse(game)
	game_response.Previews = preview_responses

	return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(
		DTO.GameResponseDTO{Game: game_response}, ""),
	)
}

// GetGames godoc
//
// @Description  get all game
// @Tags         Game
// @Produce      json
// @Success      200 {object} api.SuccessResponse[DTO.GamesResponseDTO]
// @Failure      500
// @Router       /api/v1/games/ [get]
func (h *GameHandler) GetGames(c *fiber.Ctx) error {
	games, err := h.Repository.GetGameAll()
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found game")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	games_response := h.Converter.GetGamesToGamesResponseOnly(games)

	return c.Status(fiber.StatusOK).JSON(api.NewSuccessResponse(
		DTO.GamesResponseDTO{Games: games_response}, ""),
	)
}

// CreateGame godoc
//
// @Description  create game
// @Tags         Game
// @Accept		 json
// @Produce		 json
// @Param        CreateGameInput body  DTO.CreateGameInput  true  "CreateGameInput"
// @Success		 201 {object} api.SuccessResponse[DTO.GameResponseOnlyDTO]
// @Failure      400 {object} api.ErrorResponse
// @Failure      409 {object} api.ErrorResponse
// @Failure      500 {object} api.ErrorResponse
// @Failure      422 {object} api.ErrorResponse
// @Router		 /api/v1/games [post]
func (h *GameHandler) CreateGame(c *fiber.Ctx) error {
	var payload *DTO.CreateGameInput

	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}

	gameErrors := validation.ValidateStruct(payload)
	if gameErrors != nil {
		return api.BadRequestParamError(c, gameErrors)
	}

	game, err := h.Repository.CreateGame(payload.Title, payload.Description, payload.Src, "")

	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	return c.Status(fiber.StatusCreated).JSON(api.NewSuccessResponse(
		DTO.GameResponseOnlyDTO{Game: h.Converter.GetGameToGameResponseOnly(game)},
		""))
}

// UpdateGame godoc
//
// @Description  update game
// @Tags         Game
// @Accept		 json
// @Produce		 json
// @Param        id   path string true "Game ID"
// @Param        UpdateGameInput		body		DTO.UpdateGameInput		true   "UpdateGameInput"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/games/{id} [patch]
func (h *GameHandler) UpdateGame(c *fiber.Ctx) error {
	var payload *DTO.UpdateGameInput

	if err := c.BodyParser(&payload); err != nil {
		return api.UnprocessableEntityError(c, err)
	}

	gameErrors := validation.ValidateStruct(payload)
	if gameErrors != nil {
		return api.BadRequestParamError(c, gameErrors)
	}

	id := c.Params("id")

	err := h.Repository.UpdateGame(id, payload.Title, payload.Description, payload.Src)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found game")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

// UploadGameIcon godoc
//
// @Description  upload game icon
// @Tags         Game
// @Accept		 json
// @Produce		 json
// @Param        id   path string true "Game ID"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/games/{id}/icon [post]
func (h *GameHandler) UploadGameIcon(c *fiber.Ctx) error {
	id := c.Params("id")

	iconFileHeader, err := processFormFile(c, "icon", []string{".png", ".jpg"}, true)
	if iconFileHeader == nil {
		return err
	}

	iconFileReader, err := iconFileHeader.Open()
	defer iconFileReader.Close()
	if err != nil {
		return err
	}

	imageObjName := fmt.Sprintf("/games/%s/icon%s", id, filepath.Ext(iconFileHeader.Filename))
	_, err = h.Minio.PutObject(imageObjName, iconFileReader)

	if err != nil {
		return api.InternalServerError(c, err, "something went wrong")
	}

	imagePath := h.s3ObjectNameToUrl(imageObjName)

	err = h.Repository.UpdateGameIcon(id, imagePath)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found game")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

// DeleteGame godoc
//
// @Description	 delete game by id
// @Tags         Game
// @Produce		 json
// @Param        id   path string true "Game ID"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/games/{id} [delete]
func (h *GameHandler) DeleteGame(c *fiber.Ctx) error {
	gameId := c.Params("id")

	go h.deleteGameDataFromS3(gameId)

	err := h.Repository.DeleteGame(gameId)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found game")
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// CreatePreview godoc
//
// @Description	 create game preview
// @Tags         Game
// @Produce		 json
// @Success		 201 {object} api.SuccessResponse[DTO.PreviewResponseDTO]
// @Failure      400 {object} api.ErrorResponse
// @Failure      409 {object} api.ErrorResponse
// @Failure      500 {object} api.ErrorResponse
// @Failure      422 {object} api.ErrorResponse
// @Router		 /api/v1/games/ [post]
func (h *GameHandler) CreatePreview(c *fiber.Ctx) error {
	gameId := c.FormValue("gameId")

	if gameId == "" {
		return api.BadRequestParamError(c, []*api.Error{
			{Code: api.IncorrectParameter, Parameter: "gameId", Message: "gameId is empty"},
		})
	}

	randomCode := utils.GenerateCode(10)

	// Process Image
	imageFileHeader, err := processFormFile(c, "image", []string{".png", ".jpg"}, true)
	if imageFileHeader == nil {
		return err
	}
	imageObjName := fmt.Sprintf("/games/%s/preview/%s", gameId, randomCode+"_image"+filepath.Ext(imageFileHeader.Filename))

	// Process Video
	videoFileHeader, err := processFormFile(c, "video", []string{".mp4"}, false)
	var videoObjName *string
	if videoFileHeader != nil {
		path := fmt.Sprintf("/games/%s/preview/%s", gameId, randomCode+"_video"+filepath.Ext(videoFileHeader.Filename))
		videoObjName = &path
	}

	imageCreatePath := h.s3ObjectNameToUrl(imageObjName)

	var videoCreatePath *string
	if videoFileHeader != nil {
		path := h.s3ObjectNameToUrl(*videoObjName)
		videoCreatePath = &path
	}

	// Start transaction
	preview, err := h.Repository.CreatePreviewWithTransaction(gameId, imageCreatePath, videoCreatePath, func(preview *repository.GetPreview) error {
		// Put image
		imageFileReader, err := imageFileHeader.Open()
		defer imageFileReader.Close()
		if err != nil {
			return err
		}

		_, err = h.Minio.PutObject(imageObjName, imageFileReader)

		if err != nil {
			return err
		}

		// Put video
		if videoFileHeader != nil {
			videoFileReader, err := videoFileHeader.Open()
			defer videoFileReader.Close()
			if err != nil {
				return err
			}
			_, err = h.Minio.PutObject(*videoObjName, videoFileReader)
		}

		return err
	})

	if err != nil {
		if errors.Is(err, repository.ErrForeignKeyViolation) {
			return api.NotFoundError(c, fmt.Sprintf("game with id %s doesn't exist", gameId))
		} else {
			return api.InternalServerError(c, err, "something went wrong")
		}
	}

	preview_response := h.Converter.GetPreviewToPreviewResponse(preview)

	return c.Status(fiber.StatusCreated).JSON(api.NewSuccessResponse(
		DTO.PreviewResponseDTO{Preview: preview_response},
		""))
}

// DeletePreview godoc
//
// @Description	 delete preview by id
// @Tags         Game
// @Produce		 json
// @Param        id   path string true "Preview ID"
// @Success		 200
// @Failure      500
// @Router		 /api/v1/games/preview/ [delete]
func (h *GameHandler) DeletePreview(c *fiber.Ctx) error {
	previewId := c.Params("id")

	preview, err := h.Repository.GetPreviewByID(previewId)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found preview")
		} else {
			return api.InternalServerError(c, err, "couldn't remove resources from storage")
		}
	}

	var p_video_name *string = nil
	if preview.Video != nil {
		video_name := h.getBucketNameByPath(*preview.Video)
		p_video_name = &video_name
	}
	err = h.deletePreviewFromS3(h.getBucketNameByPath(preview.Image), p_video_name)
	if err != nil {
		return api.InternalServerError(c, err, "couldn't remove preview resources from storage")
	}

	err = h.Repository.DeletePreview(previewId)

	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return api.NotFoundError(c, "not found preview")
		} else {
			return api.InternalServerError(c, err, "couldn't remove preivew")
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func processFormFile(c *fiber.Ctx, name string, availableFormtas []string, require bool) (fileHeader *multipart.FileHeader, err error) {
	fileHeader, err = c.FormFile(name)

	if fileHeader == nil {
		slog.Error(err.Error())
		if require {
			return nil, c.Status(fiber.StatusBadRequest).JSON(api.NewErrorResponse([]*api.Error{
				{Code: api.IncorrectParameter, Parameter: name, Message: "file is nil"},
			}))
		} else {
			return nil, nil
		}
	}

	ext := filepath.Ext(fileHeader.Filename)
	ext = strings.ToLower(ext)
	if !slices.Contains(availableFormtas, ext) {
		return nil, c.Status(fiber.StatusUnprocessableEntity).JSON(api.NewErrorResponse([]*api.Error{
			{Code: api.IncorrectParameter, Parameter: name, Message: "Available formats: " + strings.Join(availableFormtas, " ")},
		}))
	}

	return fileHeader, nil
}

func (h *GameHandler) deletePreviewFromS3(imagePath string, videoPath *string) error {
	err := h.Minio.RemoveObject(imagePath)

	if err != nil {
		slog.Error(err.Error())
		return err
	}
	if videoPath != nil {
		err = h.Minio.RemoveObject(*videoPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *GameHandler) deleteGameDataFromS3(id string) {
	h.Minio.RemoveFolder("games/" + id + "/")
}

func (h *GameHandler) getBucketNameByPath(path string) string {
	return strings.TrimLeft(path, h.Config.MinioOrigin+h.Config.AppBucket)
}

func (h *GameHandler) s3ObjectNameToUrl(name string) string {
	return h.Config.MinioOrigin + "/" + h.Config.AppBucket + name
}
