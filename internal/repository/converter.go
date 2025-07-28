package repository

import (
	"gameplatform/internal/database"
	"time"

	"github.com/google/uuid"
)

func TimeToTime(t time.Time) time.Time {
	return t
}

func UuidToString(id uuid.UUID) string {
	return id.String()
}

func ByteToString(b []byte) string {
	return string(b)
}

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:output:file ./generated.go
// goverter:output:format struct
// goverter:extend UuidToString TimeToTime ByteToString
type UserConverter interface {
	CreateUserToCreateUserParams(source *CreateUser) database.CreateUserParams
	PlatformUserToGetUser(source *database.PlatformUser) GetUser
	PlatformUsersToGetUsers(source []database.PlatformUser) []GetUser
}

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:output:file ./generated.go
// goverter:output:format struct
// goverter:extend UuidToString TimeToTime ByteToString
type GameConverter interface {
	PlatformGameToGetGame(source *database.PlatformGame) GetGame
	PlatformPreviewToGetPreview(source *database.PlatformPreview) GetPreview
	PlatformGamesToGetGames(platform_games []database.PlatformGame) []GetGame
	PlatformPreviewsToGetPreviews(platform_preview []database.PlatformPreview) []GetPreview
}
