package orders

import (
	"context"
	"database/sql"
	"log"

	"github.com/chris-ramon/gql-demo/internal/models"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) FindOrders(ctx context.Context, currentUser *models.User) ([]*models.Order, error) {
	orders, err := currentUser.Orders().All(ctx, repo.db)
	if err != nil {
		log.Printf("failed to find orders: %v", err)
		return nil, nil
	}
	return orders, nil
}

func NewRepository(db *sql.DB) *Repository {
	repo := Repository{db}

	return &repo
}
