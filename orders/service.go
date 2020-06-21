package orders

import (
	"context"

	"github.com/chris-ramon/gql-demo/models"
)

type Service interface {
	FindOrders(ctx context.Context, currentUser *models.User) ([]*models.Order, error)
}

type service struct {
	repo *Repository
}

func (srv *service) FindOrders(ctx context.Context, currentUser *models.User) ([]*models.Order, error) {
	return srv.repo.FindOrders(ctx, currentUser)
}

func NewService(repo *Repository) Service {
	srv := service{repo}

	return &srv
}
