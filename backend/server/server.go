package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aneri/chat-go-graphql/backend/dal"

	"github.com/gorilla/mux"

	"github.com/99designs/gqlgen/handler"
	"github.com/aneri/chat-go-graphql/backend"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db, err := dal.DbConnect()
	fmt.Println(db, err)
	router := mux.NewRouter()
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
