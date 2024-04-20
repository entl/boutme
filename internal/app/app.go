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
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
)

type App struct {
	echo          *echo.Echo
	healthHandler *endpoint.HealthHandler

	projectHandler *endpoint.ProjectHandler
	projectService *service.ProjectService
	projectRepo    *repository.SqlProjectRepository

	userHandler *endpoint.UserHandler
	userRepo    *repository.SqlUserRepository
	userService *service.UserService

	jwtHandler *endpoint.JWTHandler
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

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
	if frontendOrigin == "" {
		log.Fatal("FRONTEND_ORIGIN is not set")
	}

	frontendUrl := os.Getenv("FRONTEND_URL")
	if frontendUrl == "" {
		log.Fatal("FRONTEND_URL is not set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	db := database.New(conn)

	a := &App{}
	a.echo = echo.New()
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontendOrigin, frontendUrl},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	a.echo.Validator = &CustomValidator{validator: validator.New()}

	a.projectRepo = repository.NewSqlProjectRepository(db)
	a.userRepo = repository.NewSqlUserRepository(db)

	a.projectService = service.NewProjectService(a.projectRepo)
	a.userService = service.NewUserService(a.userRepo)

	a.healthHandler = endpoint.NewHealthHandler()
	a.projectHandler = endpoint.NewProjectHandler(a.projectService)
	a.userHandler = endpoint.NewUserHandler(a.userService)

	a.jwtHandler = endpoint.NewJWTHandler()

	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	a.echo.GET("/status", a.healthHandler.Health)

	a.echo.POST("/login", a.userHandler.AuthenticateUser)
	a.echo.POST("/register", a.userHandler.RegisterUser)

	a.echo.GET("/projects", a.projectHandler.GetProjects)
	a.echo.GET("/projects/:id", a.projectHandler.GetProject)

	config := echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	adminGroup := a.echo.Group("/admin")
	adminGroup.Use(echojwt.WithConfig(config))
	adminGroup.POST("/projects", a.projectHandler.AddProject)
	adminGroup.PUT("/projects/:id", a.projectHandler.UpdateProject)
	adminGroup.DELETE("/projects/:id", a.projectHandler.DeleteProject)
	adminGroup.GET("", a.jwtHandler.VerifyJWT)

	return a, nil
}

func (a *App) Run() error {
	hostString := os.Getenv("HOST")
	portString := os.Getenv("PORT")

	fmt.Println("server is running...")

	err := a.echo.Start(fmt.Sprintf("%s:%s", hostString, portString))
	if err != nil {
		log.Fatal("failed to start server", err)
	}

	return nil
}
