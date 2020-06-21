package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/chris-ramon/gql-demo/graphql"
	"github.com/chris-ramon/gql-demo/graphql/generated"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gql_demo_dev")
	if err != nil {
		log.Fatal(err)
	}

	boil.DebugMode = true

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphql.NewSchemaConfig(db)))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("server running on port :8080")
	log.Println("graphql playground running on :8080/playground")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
