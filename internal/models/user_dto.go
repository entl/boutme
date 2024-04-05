package models

import "github.com/google/uuid"

type UserLoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserJWTResponseDTO struct {
	Token string `json:"token"`
}

type UserCreateRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCreateResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
