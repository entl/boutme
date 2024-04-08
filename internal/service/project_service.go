package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/entl/boutme/internal/database"
	"github.com/entl/boutme/internal/models"
	"github.com/entl/boutme/internal/repository"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"time"
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
	projectParams.ID = uuid.New()
	projectParams.CreatedAt = time.Now()
	projectParams.UpdatedAt = time.Now()
	
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

func (s *ProjectService) GetProject(ctx context.Context, id string) (models.ProjectResponseDTO, error) {
	projectDb, err := s.projectRepository.GetProject(ctx, id)
	if err != nil {
		fmt.Println("error getting project service: ", err)
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

func (s *ProjectService) UpdateProject(ctx context.Context, id string, project models.ProjectUpdateDTO) (models.ProjectResponseDTO, error) {
	var projectParams database.UpdateProjectParams
	err := mapstructure.Decode(project, &projectParams)
	if err != nil {
		fmt.Println("error decoding project: ", err)
		return models.ProjectResponseDTO{}, errors.New(fmt.Sprintf("error decoding project: %v", err))
	}

	projectParams.ID, err = uuid.Parse(id)
	if err != nil {
		fmt.Println("error parsing project id: ", err)
		return models.ProjectResponseDTO{}, errors.New(fmt.Sprintf("error parsing project id: %v", err))

	}

	projectDb, err := s.projectRepository.UpdateProject(ctx, projectParams)
	if err != nil {
		fmt.Println("error updating project service: ", err)
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

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	projectId, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("error parsing project id: ", err)
		return errors.New(fmt.Sprintf("error parsing project id: %v", err))
	}

	err = s.projectRepository.DeleteProject(ctx, projectId)
	if err != nil {
		fmt.Println("error deleting project service: ", err)
		return err
	}

	return nil
}
