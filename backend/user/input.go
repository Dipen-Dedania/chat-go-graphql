package user

import (
	"context"
	"sync"

	"github.com/aneri/chat-go-graphql/backend/dal"
)

var user map[string]chan User
var mu sync.Mutex

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) JoinUser(ctx context.Context, name string) (User, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	users := User{
		Name: name,
	}
	if _, err := crConn.Db.Exec("INSERT INTO chatApp.userdata (name,createdat) VALUES ($1,NOW())", users.Name); err != nil {
		return User{}, nil
	}
	for _, observe := range user {
		observe <- users
	}
	return users, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) UserJoined(ctx context.Context) (<-chan User, error) {
	id := randString(8)
	events := make(chan User, 1)
	go func() {
		<-ctx.Done()
		delete(user, id)
	}()
	user[id] = events
	return events, nil
}
