package main

import (
	"log"
	_ "net/http"
	"os"
	"database/sql"

	_ "github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/matthewmazzanti/wordgame/srv/graph"
	_ "github.com/matthewmazzanti/wordgame/srv/graph/generated"
	_ "github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/wordgame")
	if err != nil {
		log.Fatal(err)
	}

	mask, primary := makeMasks("abcdefg")
	rows, err := db.Query(
		"select word from word where bitmap & ? = 0 and bitmap & ? > 0;",
		mask,
		primary,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		var word string;

		err := rows.Scan(&word)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(word)
		count++
	}

	print(count)

	/*
	resolver := graph.Resolver{ DB: db }
	config := generated.Config{ Resolvers: &resolver }
	schema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(schema)

	router := chi.NewRouter()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	*/
}

func makeMasks(letters string) (uint32, uint32) {
	var mask uint32 = 0
	for i := 1; i < len(letters); i++ {
		letter := letters[i]
		mask = mask | 1 << (int(letter) - 97)
	}

	var primary uint32 = 1 << (int(letters[0]) - 97)
	mask = ^(mask | primary)

	return mask, primary
}
