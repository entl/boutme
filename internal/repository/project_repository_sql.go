package repository

import (
	"context"
	"fmt"
	"github.com/entl/boutme/internal/database"
	"github.com/google/uuid"
)

type SqlProjectRepository struct {
	db *database.Queries
}

func NewSqlProjectRepository(db *database.Queries) *SqlProjectRepository {
	return &SqlProjectRepository{db: db}
}

func (repository *SqlProjectRepository) CreateProject(ctx context.Context, projectParams database.CreateProjectParams) (database.Project, error) {
	projectParams.ID = uuid.New()
	projectDB, err := repository.db.CreateProject(ctx, projectParams)

	if err != nil {
		fmt.Println("error creating project rep: ", err)
		return database.Project{}, err
	}

	return projectDB, nil
}

func (repository *SqlProjectRepository) GetProjects(ctx context.Context) ([]database.Project, error) {
	projects, err := repository.db.GetProjects(ctx)
	if err != nil {
		fmt.Println("error getting projects rep: ", err)
		return []database.Project{}, err
	}

	return projects, nil
}

func (repository *SqlProjectRepository) GetProject(ctx context.Context, id string) (database.Project, error) {
	projectId, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("error parsing project id: ", err)
		return database.Project{}, err
	}

	project, err := repository.db.GetProjectByID(ctx, projectId)
	if err != nil {
		fmt.Println("error getting project rep: ", err)
		return database.Project{}, err
	}

	return project, nil
}

func (repository *SqlProjectRepository) UpdateProject(ctx context.Context, projectParams database.UpdateProjectParams) (database.Project, error) {
	project, err := repository.db.UpdateProject(ctx, projectParams)
	if err != nil {
		fmt.Println("error updating project rep: ", err)
		return database.Project{}, err
	}

	return project, nil
}

func (repository *SqlProjectRepository) DeleteProject(ctx context.Context, projectId uuid.UUID) error {
	err := repository.db.DeleteProject(ctx, projectId)
	if err != nil {
		fmt.Println("error deleting project rep: ", err)
		return err
	}

	return nil
}
