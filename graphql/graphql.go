package graphql

import (
	"context"
	"time"

	"github.com/volatiletech/null"

	"github.com/chris-ramon/gql-demo/chats"
	"github.com/chris-ramon/gql-demo/graphql/generated"
	"github.com/chris-ramon/gql-demo/models"
	"github.com/chris-ramon/gql-demo/orders"
	"github.com/chris-ramon/gql-demo/users"
)

type RootResolver struct {
	us users.Service
	os orders.Service
	cs chats.Service
}

func (r *RootResolver) Query() generated.QueryResolver {
	return &queryResolver{
		RootResolver: r,
		us:           r.us}
}

func (r *RootResolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{
		RootResolver: r,
		cs:           r.cs,
	}
}

func (r *RootResolver) User() generated.UserResolver {
	return &userResolver{
		RootResolver: r,
		os:           r.os,
	}
}

func (r *RootResolver) Chat() generated.ChatResolver {
	return &chatResolver{r}
}

type queryResolver struct {
	*RootResolver
	us users.Service
}

type userResolver struct {
	*RootResolver
	os orders.Service
}

type chatResolver struct {
	*RootResolver
}

func (r *RootResolver) TotalUnreadMessages(ctx context.Context, obj *models.Chat) (*int, error) {
  return &obj.TotalUnreadMessages.Int, nil
}

type subscriptionResolver struct {
	*RootResolver
	cs chats.Service
}

func (sr *subscriptionResolver) Chats(ctx context.Context) (<-chan []*models.Chat, error) {
	chats := make(chan []*models.Chat, 1)
	c1 := models.Chat{UUID: "345", TotalUnreadMessages: null.Int{Int: 0, Valid: true}}
	c2 := models.Chat{UUID: "987", TotalUnreadMessages: null.Int{Int: 0, Valid: true}}
	go func() {
		for {
			time.Sleep(1 * time.Second)
			var newChats []*models.Chat
			c1.TotalUnreadMessages = null.Int{Int: c1.TotalUnreadMessages.Int + 1, Valid: true}
			newChats = append(newChats, &c1)
			newChats = append(newChats, &c2)
			chats <- newChats
		}
	}()
	return chats, nil
}

func (qr *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	return qr.us.FindUser(ctx, 1)
}

func (ur *userResolver) Orders(ctx context.Context, currentUser *models.User) ([]*models.Order, error) {
	return ur.os.FindOrders(ctx, currentUser)
}

func NewRootResolver(us users.Service, os orders.Service, cs chats.Service) *RootResolver {
	return &RootResolver{us, os, cs}
}

func NewSchemaConfig(us users.Service, os orders.Service, cs chats.Service) generated.Config {
	return generated.Config{
		Resolvers: NewRootResolver(us, os, cs),
	}
}
