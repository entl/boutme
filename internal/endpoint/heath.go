package endpoint

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(ctx echo.Context) error {
	err := ctx.JSON(http.StatusOK, "Server healthy")
	if err != nil {
		return err
	}

	return nil
}
