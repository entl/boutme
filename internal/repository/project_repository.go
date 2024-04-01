package repository

import (
	"context"
	"github.com/entl/boutme/internal/database"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project database.CreateProjectParams) (*database.Project, error)
	GetProjects(ctx context.Context) ([]*database.Project, error)
	UpdateProject(ctx context.Context, project database.UpdateProjectParams) (*database.Project, error)
	DeleteProject(ctx context.Context, id string) error
}
