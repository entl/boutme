package endpoint

import (
	"fmt"
	"github.com/entl/boutme/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type JWTHandler struct {
	authService *service.AuthService
}

func NewJWTHandler(authService *service.AuthService) *JWTHandler {
	return &JWTHandler{authService: authService}
}

func (h *JWTHandler) VerifyJWT(ctx echo.Context) error {
	// change later
	return ctx.JSON(http.StatusOK, map[string]string{"message": "jwt valid"})
}

func (h *JWTHandler) RefreshJWT(ctx echo.Context) error {
	newToken, err := h.authService.IssueJWT()
	if err != nil {
		fmt.Println("Error refreshing token", err)
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"token": newToken})
}
