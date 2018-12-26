package dal

import (
	"context"
	"log"
	"net/http"
)

var ctxt context.Context

// DbMiddleWareHandler middleware for database
func DbMiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn, err := DbConnect()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(ctxt))
	})
}
