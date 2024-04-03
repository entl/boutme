package endpoint

import (
	"fmt"
	"github.com/entl/boutme/internal/models"
	"github.com/entl/boutme/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) AddProject(ctx echo.Context) error {
	var project models.ProjectRequestDTO
	if err := ctx.Bind(&project); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "cannot bind")
	}
	if err := ctx.Validate(project); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("error validating project: %v", err))
	}

	projectRes, err := h.projectService.AddProject(ctx.Request().Context(), project)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error creating project handler")
	}
	return ctx.JSON(http.StatusCreated, projectRes)
}

func (h *ProjectHandler) GetProjects(ctx echo.Context) error {
	projectRes, err := h.projectService.GetProjects(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting projects handler")
	}

	return ctx.JSON(http.StatusOK, projectRes)
}

func (h *ProjectHandler) GetProject(ctx echo.Context) error {
	id := ctx.Param("id")
	projectRes, err := h.projectService.GetProject(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error getting project handler")
	}

	return ctx.JSON(http.StatusOK, projectRes)
}

func (h *ProjectHandler) UpdateProject(ctx echo.Context) error {
	id := ctx.Param("id")

	var project models.ProjectUpdateDTO
	if err := ctx.Bind(&project); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "cannot bind")
	}
	if err := ctx.Validate(project); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("error validating project: %v", err))
	}

	projectRes, err := h.projectService.UpdateProject(ctx.Request().Context(), id, project)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error updating project handler")
	}
	return ctx.JSON(http.StatusOK, projectRes)
}

func (h *ProjectHandler) DeleteProject(ctx echo.Context) error {
	id := ctx.Param("id")

	err := h.projectService.DeleteProject(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error deleting project handler")
	}
	return ctx.NoContent(http.StatusNoContent)
}
