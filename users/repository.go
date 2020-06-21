package users

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/chris-ramon/gql-demo/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	repo := Repository{db}
	return &repo
}

func (repo *Repository) FindUser(ctx context.Context, ID int) (*models.User, error) {
	user, err := models.FindUser(ctx, repo.db, ID)
	if err != nil {
		log.Printf("failed to find user: %v", err)
		return nil, errors.New("user not found")
	}
	return user, nil
}
