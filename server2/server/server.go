package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/aneri/chat-go-graphql/server2"
	server "github.com/aneri/chat-go-graphql/server2"
	"github.com/gorilla/mux"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := mux.NewRouter()
	router.Use(server.MiddleWareHandler)
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(server2.NewExecutableSchema(server2.Config{Resolvers: &server2.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
