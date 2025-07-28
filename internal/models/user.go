package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                 *uuid.UUID
	Name               string
	Email              string
	Password           []byte
	IsAdmin            bool
	VerificationCode   string
	Verified           bool
	Birthday           *time.Time
	Gender             string
	Grade              int
	ContinuousProgress int
	CreatedAt          *time.Time
}
