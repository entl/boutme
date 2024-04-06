package endpoint

import (
	"github.com/entl/boutme/internal/models"
	"github.com/entl/boutme/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) AuthenticateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.userService.AuthenticateUser(c.Request().Context(), username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "error authenticating user")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	var user models.UserCreateRequestDTO

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "cannot bind")
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "error validating user")
	}

	userRes, err := h.userService.RegisterUser(c.Request().Context(), user.Username, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error registering user")
	}

	return c.JSON(http.StatusCreated, userRes)
}
