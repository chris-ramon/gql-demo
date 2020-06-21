package graphql

import (
	"context"

	"github.com/chris-ramon/gql-demo/graphql/generated"
	"github.com/chris-ramon/gql-demo/models"
	"github.com/chris-ramon/gql-demo/orders"
	"github.com/chris-ramon/gql-demo/users"
)

type Resolver struct {
	us users.Service
	os orders.Service
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{
    Resolver: r,
    us: r.us}
}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{
    Resolver: r,
    os: r.os,
  }
}

type queryResolver struct{
  *Resolver
	us users.Service
}
type userResolver struct{
  *Resolver
	os orders.Service
}

func (qr *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	return qr.us.FindUser(ctx, 1)
}

func (ur *userResolver) Orders(ctx context.Context, currentUser *models.User) ([]*models.Order, error) {
	return ur.os.FindOrders(ctx, currentUser)
}

func NewResolver(us users.Service, os orders.Service) *Resolver {
	return &Resolver{us, os}
}

func NewSchemaConfig(us users.Service, os orders.Service) generated.Config {
	return generated.Config{
		Resolvers: NewResolver(us, os),
	}
}
