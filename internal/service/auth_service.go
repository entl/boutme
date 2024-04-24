package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) IssueJWT() (string, error) {
	claims := jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60))}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("error signing token: ", err)
		return "", err
	}

	return tokenString, nil
}
