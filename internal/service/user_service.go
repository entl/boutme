package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/entl/boutme/internal/database"
	"github.com/entl/boutme/internal/models"
	"github.com/entl/boutme/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	userRepository *repository.SqlUserRepository
}

func NewUserService(userRepository *repository.SqlUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) AuthenticateUser(ctx context.Context, username string, password string) (models.UserJWTResponseDTO, error) {
	userDB, err := s.userRepository.GetUser(ctx, username)
	if err != nil {
		fmt.Println("error getting user: ", err)
		return models.UserJWTResponseDTO{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(password))
	if err != nil {
		fmt.Println("error comparing password: ", err)
		return models.UserJWTResponseDTO{}, err
	}

	claims := jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60))}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println("error signing tokcen: ", err)
		return models.UserJWTResponseDTO{}, err
	}

	return models.UserJWTResponseDTO{Token: tokenString}, nil
}

func (s *UserService) RegisterUser(ctx context.Context, username string, password string) (models.UserCreateResponseDTO, error) {
	numUsers, err := s.userRepository.CountUsers(ctx)
	if err != nil {
		fmt.Println("error counting users: ", err)
		return models.UserCreateResponseDTO{}, err
	}
	if numUsers > 0 {
		fmt.Println("only one user allowed")
		return models.UserCreateResponseDTO{}, errors.New("only one user allowed")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error hashing password: ", err)
		return models.UserCreateResponseDTO{}, err
	}

	user, err := s.userRepository.AddUser(ctx, database.AddUserParams{
		ID:        uuid.New(),
		Username:  username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		fmt.Println("error adding user: ", err)
		return models.UserCreateResponseDTO{}, err
	}

	return models.UserCreateResponseDTO{ID: user.ID, Username: user.Username}, nil
}
