package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: there no one suitable note")

type Game struct {
	ID           int
	Title        string
	Description  string
	Src          string
	PreviewImage string
	Created      time.Time
}
