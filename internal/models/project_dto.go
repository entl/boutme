package models

import (
	"github.com/google/uuid"
)

type ProjectRequestDTO struct {
	Name        string   `json:"name" mapstructure:"Name" validate:"required"`
	Description string   `json:"description" mapstructure:"Description" validate:"required"`
	ImageUrl    string   `json:"image_url" mapstructure:"ImageUrl" validate:"required"`
	GithubUrl   string   `json:"github_url" mapstructure:"GithubUrl" validate:"required"`
	TechStack   []string `json:"tech_stack" mapstructure:"TechStack" validate:"required"`
}

type ProjectResponseDTO struct {
	ID          uuid.UUID `json:"id" mapstructure:"ID" validate:"required"`
	Name        string    `json:"name" mapstructure:"Name" validate:"required"`
	Description string    `json:"description" mapstructure:"Description" validate:"required"`
	ImageUrl    string    `json:"image_url" mapstructure:"ImageUrl" validate:"required"`
	GithubUrl   string    `json:"github_url" mapstructure:"GithubUrl" validate:"required"`
	TechStack   []string  `json:"tech_stack" mapstructure:"TechStack" validate:"required"`
}
