package repository

import (
	"context"
	"github.com/entl/boutme/internal/database"
)

type SqlUserRepository struct {
	db  *database.Queries
	ctx context.Context
}

func NewSqlUserRepository(db *database.Queries) *SqlUserRepository {
	return &SqlUserRepository{db: db}
}

func (repository *SqlUserRepository) GetUser(username string) (database.User, error) {
	return repository.db.GetUser(repository.ctx, username)
}
