package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aneri/chat-go-graphql/backend/chatConversation"
	"github.com/aneri/chat-go-graphql/backend/dal"
	"github.com/aneri/chat-go-graphql/backend/graph"
	"github.com/aneri/chat-go-graphql/backend/user"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/99designs/gqlgen/handler"
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
	router.Handle("/", handler.Playground("GraphQL playground", "/user"))

	router.Handle("/user", corsAccess(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &user.Resolver{}}),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(request *http.Request) bool {
				return true
			},
		}),
	)))
	router.Handle("/chat", corsAccess(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &chatConversation.Resolver{}}),
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
