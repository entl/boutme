package repository

import (
	"context"
	"github.com/entl/boutme/internal/database"
)

type UserRepository interface {
	GetUser(ctx context.Context) (*database.User, error)
}
