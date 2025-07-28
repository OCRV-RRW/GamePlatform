package repository

import (
	"context"
	"gameplatform/internal/database"
	"gameplatform/internal/dbconn"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetGame struct {
	ID          uuid.UUID
	Title       string
	Description string
	Src         string
	Icon        string
}

type GetPreview struct {
	ID    uuid.UUID
	Image string
	Video *string
}

// func PlatformGameToGetGame(game database.PlatformGame) GetGame {
// 	return GetGame{
// 		ID:          game.ID,
// 		Title:       game.Title,
// 		Description: game.Description,
// 		Src:         game.Src,
// 		Icon:        game.Icon,
// 	}
// }
//
// func PlatformPreviewToGetPreview(preview database.PlatformPreview) GetPreview {
// 	return GetPreview{
// 		ID:    preview.ID,
// 		Image: preview.Image,
// 		Video: preview.Video,
// 	}
// }

// func PlatformGamesToGetGames(platform_games []database.PlatformGame) []GetGame {
// 	length := len(platform_games)
// 	games := make([]GetGame, length)
// 	for i := 0; i < length; i++ {
// 		games[i] = PlatformGameToGetGame(platform_games[i])
// 	}
//
// 	return games
// }
//
// func PlatformPreviewsToGetPreviews(platform_preview []database.PlatformPreview) []GetPreview {
// 	length := len(platform_preview)
// 	previews := make([]GetPreview, length)
// 	for i := 0; i < length; i++ {
// 		previews[i] = PlatformPreviewToGetPreview(platform_preview[i])
// 	}
//
// 	return previews
// }

type GameRepository struct {
	db   *dbconn.DatabaseConnection
	conv GameConverter
}

func NewGameRepository(db *dbconn.DatabaseConnection, conv GameConverter) *GameRepository {
	return &GameRepository{db: db, conv: conv}
}

func (r *GameRepository) GetGameByID(id string) (*GetGame, error) {
	ctx := context.Background()
	game_id, err := uuid.Parse(id)
	if err != nil {
		return nil, Error
	}

	platform_game, err := r.db.Queries.GetGameByID(ctx, game_id)
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}

	result := r.conv.PlatformGameToGetGame(&platform_game)

	return &result, nil
}

func (r *GameRepository) GetGameAll() ([]GetGame, error) {
	ctx := context.Background()
	platform_games, err := r.db.Queries.GetAllGames(ctx)
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}

	return r.conv.PlatformGamesToGetGames(platform_games), nil
}

func (r *GameRepository) CreateGame(title string, description string, src string, icon string) (*GetGame, error) {
	ctx := context.Background()
	game, err := r.db.Queries.CreateGame(ctx, database.CreateGameParams{
		Title:       title,
		Description: description,
		Src:         src,
		Icon:        icon,
	})
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}
	result := r.conv.PlatformGameToGetGame(&game)

	return &result, nil
}

func (r *GameRepository) UpdateGame(id string, title string, description string, scr string, icon string) error {
	ctx := context.Background()
	game_id, err := uuid.Parse(id)
	if err != nil {
		return SqlcErrToRepositoryErr(err)
	}

	err = r.db.Queries.UpdateGame(ctx, database.UpdateGameParams{
		ID:          game_id,
		Title:       title,
		Description: description,
		Src:         scr,
		Icon:        icon,
	})

	return SqlcErrToRepositoryErr(err)
}

func (r *GameRepository) DeleteGame(id string) error {
	ctx := context.Background()
	game_id, err := uuid.Parse(id)
	if err != nil {
		return SqlcErrToRepositoryErr(err)
	}
	err = r.db.Queries.DeleteGame(ctx, game_id)

	return SqlcErrToRepositoryErr(err)
}

func (r *GameRepository) GetPreviewByID(id string) (*GetPreview, error) {
	ctx := context.Background()
	previewId, err := uuid.Parse(id)
	if err != nil {
		return nil, Error
	}

	platform_preview, err := r.db.Queries.GetPreviewByID(ctx, previewId)
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}

	preivew := r.conv.PlatformPreviewToGetPreview(&platform_preview)

	return &preivew, nil
}

func (r *GameRepository) GetPreviews(gameId string) ([]GetPreview, error) {
	ctx := context.Background()
	id, err := uuid.Parse(gameId)
	if err != nil {
		return nil, Error
	}

	platform_previews, err := r.db.Queries.GetGamePreview(ctx, id)
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}
	result := r.conv.PlatformPreviewsToGetPreviews(platform_previews)

	return result, nil
}

func createPreview(queries *database.Queries, gameId string, image string, video *string) (*GetPreview, error) {
	ctx := context.Background()
	id, err := uuid.Parse(gameId)
	if err != nil {
		return nil, Error
	}

	preview, err := queries.CreatePreview(ctx, database.CreatePreviewParams{
		GameID: id,
		Image:  image,
		Video:  video,
	})
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}

	get_preview := GetPreview{
		ID:    preview.ID,
		Image: preview.Image,
		Video: preview.Video,
	}

	return &get_preview, nil
}

func (r *GameRepository) CreatePreview(gameId string, image string, video *string) (*GetPreview, error) {
	return createPreview(r.db.Queries, gameId, image, video)
}

func (r *GameRepository) CreatePreviewWithTransaction(gameId string, image string, video *string, callback func(*GetPreview) error) (*GetPreview, error) {
	ctx := context.Background()
	tx, err := r.db.Connection.BeginTx(ctx, pgx.TxOptions{})
	tx_queries := r.db.Queries.WithTx(tx)

	preview, err := createPreview(tx_queries, gameId, image, video)
	if err != nil {
		return nil, SqlcErrToRepositoryErr(err)
	}
	err = callback(preview)
	if err == nil {
		tx.Commit(ctx)
	} else {
		tx.Rollback(ctx)
	}

	return preview, err
}

func (r *GameRepository) DeletePreview(id string) error {
	ctx := context.Background()
	previewId, err := uuid.Parse(id)
	if err != nil {
		return Error
	}

	err = r.db.Queries.DeletePreview(ctx, previewId)

	return SqlcErrToRepositoryErr(err)
}
