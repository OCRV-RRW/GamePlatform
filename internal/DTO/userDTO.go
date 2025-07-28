package DTO

import (
	"github.com/google/uuid"
	"time"
)

type TokenResponse struct {
	AccessToken string     `json:"access_token"`
	ExpiredIn   *time.Time `json:"expired_in"`
}

type UserResponseDTO struct {
	User UserResponse `json:"user"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"is_admin"`
	Birthday  *time.Time `json:"birthday"`
	Gender    string     `json:"gender"`
	CreatedAt time.Time  `json:"created_at"`
}

type UsersResponse struct {
	Users []UserResponse `json:"users"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type UpdateUserInput struct {
	Name     *string    `json:"name,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Gender   *string    `json:"gender,omitempty"`
	IsAdmin  *bool      `json:"is_admin,omitempty"`
}
