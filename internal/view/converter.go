package view

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
// goverter:output:format struct
// goverter:output:file ./generated.go
// goverter:extend TimeToTime UuidToString
type ViewConverter interface {
	GetGameToGameView(source *repository.GetGame) GameView
	GetGamesToGameViews(source []repository.GetGame) []GameView
	GetPreviewToPreviewView(source *repository.GetPreview) PreviewView
	GetPreviewsToPreviewViews(source []repository.GetPreview) []PreviewView
}
