package app

import (
	"database/sql"
	"fmt"
	"github.com/entl/boutme/internal/database"
	"github.com/entl/boutme/internal/endpoint"
	"github.com/entl/boutme/internal/repository"
	"github.com/entl/boutme/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

type App struct {
	echo           *echo.Echo
	healthHandler  *endpoint.HealthHandler
	projectHandler *endpoint.ProjectHandler
	projectService *service.ProjectService
	projectRepo    *repository.SqlProjectRepository
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() (*App, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	db := database.New(conn)

	a := &App{}

	a.echo = echo.New()

	a.echo.Validator = &CustomValidator{validator: validator.New()}

	a.projectRepo = repository.NewSqlProjectRepository(db)

	a.projectService = service.NewProjectService(a.projectRepo)

	a.healthHandler = endpoint.NewHealthHandler()
	a.projectHandler = endpoint.NewProjectHandler(a.projectService)

	a.echo.GET("/status", a.healthHandler.Health)
	a.echo.GET("/projects", a.projectHandler.GetProjects)
	a.echo.POST("/projects", a.projectHandler.AddProject)

	return a, nil
}

func (a *App) Run() error {

	fmt.Println("server is running...")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal("failed to start server", err)
	}

	return nil
}
