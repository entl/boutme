package endpoint

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type JWTHandler struct {
}

func NewJWTHandler() *JWTHandler {
	return &JWTHandler{}
}

func (h *JWTHandler) VerifyJWT(ctx echo.Context) error {
	// change later
	return ctx.JSON(http.StatusOK, map[string]string{"message": "jwt valid"})
}
