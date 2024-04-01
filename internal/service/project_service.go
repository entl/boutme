package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/entl/boutme/internal/database"
	"github.com/entl/boutme/internal/models"
	"github.com/entl/boutme/internal/repository"
	"github.com/mitchellh/mapstructure"
)

type ProjectService struct {
	projectRepository *repository.SqlProjectRepository
}

func NewProjectService(projectRepository *repository.SqlProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
	}
}

func (s *ProjectService) AddProject(ctx context.Context, project models.ProjectRequestDTO) (models.ProjectResponseDTO, error) {
	var projectParams database.CreateProjectParams

	err := mapstructure.Decode(project, &projectParams)

	if err != nil {
		fmt.Println("error decoding project: ", err)
		return models.ProjectResponseDTO{}, errors.New(fmt.Sprintf("error decoding project: %v", err))
	}

	projectDb, err := s.projectRepository.CreateProject(ctx, projectParams)
	if err != nil {
		fmt.Println("error creating project service: ", err)
		return models.ProjectResponseDTO{}, err
	}

	var projectResponse models.ProjectResponseDTO
	err = mapstructure.Decode(projectDb, &projectResponse)
	if err != nil {
		fmt.Println("error decoding project: ", err)
		return models.ProjectResponseDTO{}, err
	}

	return projectResponse, nil
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]models.ProjectResponseDTO, error) {
	projectDb, err := s.projectRepository.GetProjects(ctx)
	if err != nil {
		fmt.Println("error getting projects service: ", err)
		return []models.ProjectResponseDTO{}, err
	}

	var projectResponse []models.ProjectResponseDTO
	err = mapstructure.Decode(projectDb, &projectResponse)
	if err != nil {
		fmt.Println("error decoding project: ", err)
		return []models.ProjectResponseDTO{}, err
	}

	return projectResponse, nil
}
