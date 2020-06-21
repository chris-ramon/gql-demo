package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/chris-ramon/gql-demo/chats"
	"github.com/chris-ramon/gql-demo/graphql"
	"github.com/chris-ramon/gql-demo/graphql/generated"
	"github.com/chris-ramon/gql-demo/orders"
	"github.com/chris-ramon/gql-demo/users"
)

func main() {
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DB")

	connectionStr := fmt.Sprintf("%s:%s@/%s", mysqlUser, mysqlPassword, mysqlDB)
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	boil.DebugMode = true

	var (
		ur *users.Repository  = users.NewRepository(db)
		or *orders.Repository = orders.NewRepository(db)
		cr *chats.Repository  = chats.NewRepository(db)

		us users.Service  = users.NewService(ur)
		os orders.Service = orders.NewService(or)
		cs chats.Service  = chats.NewService(cr)
	)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphql.NewSchemaConfig(us, os, cs)))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Println("server running on port :8080")
	log.Println("graphql playground running on :8080/playground")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
