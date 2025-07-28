package DTO

import (
	"gameplatform/internal/repository"
	"time"

	"github.com/google/uuid"
)

func TimeToTime(t time.Time) time.Time {
	return t
}

func UuidToString(id uuid.UUID) string {
	return id.String()
}

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:output:file ./generated.go
// goverter:output:format struct
// goverter:extend TimeToTime
type UserConverter interface {
	GetUserToUserResponse(source *repository.GetUser) UserResponse
	GetUsersToUserResponses(source []repository.GetUser) []UserResponse
	GetUserToUpdateUser(source *repository.GetUser) repository.UpdateUser
}

// goverter:converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:output:file ./generated.go
// goverter:output:format struct
// goverter:extend TimeToTime UuidToString
type GameConverter interface {
	GetGameToGameResponseOnly(source *repository.GetGame) GameResponseOnly
	GetGamesToGamesResponseOnly(source []repository.GetGame) []GameResponseOnly
	// goverter:ignore Previews
	GetGameToGameResponse(source *repository.GetGame) GameResponse
	// goverter:ignore Previews
	GameResponseOnlyToGameResponse(source *GameResponseOnly) GameResponse
	GetPreviewToPreviewResponse(source *repository.GetPreview) PreviewResponse
	GetPreviewsToPreviewResponses(source []repository.GetPreview) []PreviewResponse
}
