package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/aneri/chat-go-graphql/backend"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := mux.NewRouter()
	router.Use(backend.MiddleWareHandler)
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))

	router.Handle("/query", corsAccess(handler.GraphQL(backend.NewExecutableSchema(backend.Config{Resolvers: &backend.Resolver{}}),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(request *http.Request) bool {
				return true
			},
		}),
	)))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
func corsAccess(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		response.Header().Set("Access-Control-Allow-Headers", "Accept, X-Requested-With, Content-Type, Authorization")
		next(response, request)
	})
}
