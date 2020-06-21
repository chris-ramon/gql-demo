package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chris-ramon/gql-demo/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type Resolver struct {
	db *sql.DB
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) User() UserResolver {
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

func NewSchemaConfig(db *sql.DB) Config {
	return Config{
		Resolvers: NewResolver(db),
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@/gql_demo_dev")
	if err != nil {
		log.Fatal(err)
	}

	boil.DebugMode = true

	srv := handler.NewDefaultServer(NewExecutableSchema(NewSchemaConfig(db)))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("server running on port :8080")
	log.Println("graphql playground running on :8080/playground")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
