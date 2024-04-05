package repository

import (
	"context"
	"fmt"
	"github.com/entl/boutme/internal/database"
)

type SqlUserRepository struct {
	db *database.Queries
}

func NewSqlUserRepository(db *database.Queries) *SqlUserRepository {
	return &SqlUserRepository{db: db}
}

func (repository *SqlUserRepository) GetUser(ctx context.Context, username string) (database.User, error) {
	user, err := repository.db.GetUser(ctx, username)
	if err != nil {
		fmt.Println("error getting user rep: ", err)
		return database.User{}, err
	}

	return user, nil
}

func (repository *SqlUserRepository) AddUser(ctx context.Context, userParams database.AddUserParams) (database.User, error) {
	user, err := repository.db.AddUser(ctx, userParams)
	if err != nil {
		fmt.Println("error adding user rep: ", err)
		return database.User{}, err
	}

	return user, nil
}

func (repository *SqlUserRepository) CountUsers(ctx context.Context) (int64, error) {
	count, err := repository.db.CountUsers(ctx)
	if err != nil {
		fmt.Println("error counting users rep: ", err)
		return 0, err
	}

	return count, nil
}
