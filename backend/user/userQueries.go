package user

import (
	"context"
	"log"

	"github.com/aneri/chat-go-graphql/backend/dal"
)

var ctxt context.Context

func GetId() string {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var user User
	row, err := crConn.Db.Query("SELECT (id) FROM training.chatuser")
	if err != nil {
		log.Fatal("error while fetching id", err)
	}
	row.Scan(&user.ID)
	defer row.Close()
	return user.ID
}
