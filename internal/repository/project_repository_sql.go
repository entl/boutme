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
