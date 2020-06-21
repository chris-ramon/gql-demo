package main

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/chris-ramon/gql-demo/models"
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
	user, err := models.FindUser(context.TODO(), qr.db, 1)
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

	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(graphql.StartOperationTrace(r.Context()))

		schema := NewExecutableSchema(NewSchemaConfig(db))
		exec := executor.New(schema)
		exec.Use(extension.Introspection{})

		var params graphql.RawParams
		start := graphql.Now()

		dec := json.NewDecoder(r.Body)
		dec.UseNumber()

		if err := dec.Decode(&params); err != nil {
			log.Printf("failed to decode request body: %v", err)
			http.Error(w, "failed 1", http.StatusBadRequest)
			return
		}

		params.ReadTime = graphql.TraceTiming{
			Start: start,
			End:   graphql.Now(),
		}

		rc, err := exec.CreateOperationContext(r.Context(), &params)
		if err != nil {
			resp := exec.DispatchError(graphql.WithOperationContext(r.Context(), rc), err)
			b, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, "failed 2", http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}

		ctx := graphql.WithOperationContext(r.Context(), rc)
		responses, ctx := exec.DispatchOperation(r.Context(), rc)

		b, err2 := json.Marshal(responses(ctx))
		if err2 != nil {
			http.Error(w, "failed 3", http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})
	log.Println("server running on port :8080")
	log.Println("graphql playground running on :8080/playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
