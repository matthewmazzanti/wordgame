package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/matthewmazzanti/wordgame/srv/graph"
	"github.com/matthewmazzanti/wordgame/srv/graph/generated"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open(
		"mysql",
		"root:password@tcp(127.0.0.1:3306)/wordgame"
	)
	if err != nil {
		log.Fatal(err)
	}

	resolver := graph.Resolver{ DB: db }
	config := generated.Config{ Resolvers: &resolver }
	schema := generated.NewExecutableSchema()
	srv := handler.NewDefaultServer(schema)

	router := chi.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
