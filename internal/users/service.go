package users

import (
	"context"

	"github.com/chris-ramon/gql-demo/internal/models"
)

type Service interface {
	FindUser(ctx context.Context, ID int) (*models.User, error)
}

type service struct {
	repo *Repository
}

func (srv *service) FindUser(ctx context.Context, ID int) (*models.User, error) {
	return srv.repo.FindUser(ctx, ID)
}

func NewService(repo *Repository) Service {
	srv := service{repo}

	return &srv
}
