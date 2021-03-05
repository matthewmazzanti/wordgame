package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"
	"math/rand"
	"time"
	_ "fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/handler/extension"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/matthewmazzanti/wordgame/srv/game"
	"github.com/matthewmazzanti/wordgame/srv/graph"
	"github.com/matthewmazzanti/wordgame/srv/graph/generated"
	"github.com/gorilla/websocket"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {
	rand.Seed(time.Now().UnixNano())
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var db *sql.DB
	for true {
		var err error
		db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/wordgame")
		if err == nil {
			break
		}

		log.Println("sleeping 5 seconds")
		time.Sleep(5 * time.Second)
	}

	resolver := graph.NewResolver(db)
	config := generated.Config{ Resolvers: resolver }
	schema := generated.NewExecutableSchema(config)
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "lambda.olympus:8080" || r.Host == "localhost:8080"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.Use(extension.Introspection{})

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{
			"http://localhost:8080",
			"http://lambda.olympus:8080",
			"http://localhost:3000",
			"http://lambda.olympus:3000",
		},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)
	router.Use(resolver.SetCookies)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
