package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/chris-ramon/gql-demo/models"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (qr *queryResolver) CurrentUser(ctx context.Context) (*models.User, error) {
	return &models.User{ID: "1", FirstName: "gopher", LastName: "gopher"}, nil
}

func (qr *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
  var orders []*models.Order
  orders = append(orders, &models.Order{ID: "2"}, &models.Order{ID: "3"})
  return orders, nil
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func NewSchemaConfig() Config {
	return Config{
		Resolvers: NewResolver(),
	}
}

func main() {
	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(graphql.StartOperationTrace(r.Context()))

		schema := NewExecutableSchema(NewSchemaConfig())
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
