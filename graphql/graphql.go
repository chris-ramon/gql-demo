package graphql

import (
	"context"
	"database/sql"
	"log"

	"github.com/chris-ramon/gql-demo/graphql/generated"
	"github.com/chris-ramon/gql-demo/models"
)

type Resolver struct {
	db *sql.DB
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

func (qr *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	user, err := models.FindUser(context.Background(), qr.db, 1)
	if err != nil {
		log.Printf("failed to find user: %v", err)
		return nil, nil
	}
	return user, nil
}

func (ur *userResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	orders, err := obj.Orders().All(ctx, ur.db)
	if err != nil {
		log.Printf("failed to find orders: %v", err)
		return orders, nil
	}

	return orders, nil
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{db}
}

func NewSchemaConfig(db *sql.DB) generated.Config {
	return generated.Config{
		Resolvers: NewResolver(db),
	}
}
