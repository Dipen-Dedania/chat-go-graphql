package backend

import (
	"context"
	"log"

	"github.com/aneri/chat-go-graphql/backend/dal"
)

var user map[string]chan User

func (r *mutationResolver) JoinUser(ctx context.Context, name string) (User, error) {
	// Database connection
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	// create User
	users := User{
		Name: name,
	}
	if _, err := crConn.Db.Exec("INSERT INTO userdata (name, createdat) VALUES ($1, NOW())", users.Name); err != nil {
		return User{}, err
	}

	// Observer new user joined
	for _, observe := range user {
		observe <- users
	}
	// return user data
	return users, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var users []User
	var user User
	row, err := crConn.Db.Query("SELECT id,name,createdat FROM userdata")
	if err != nil {
		log.Fatal("Error while retrieving user data", err)
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	defer row.Close()
	return users, nil
}
func (r *subscriptionResolver) Userjoined(ctx context.Context) (<-chan User, error) {
	var users User
	for _, observe := range user {
		observe <- users
	}
	id := randString(8)
	events := make(chan User, 1)
	go func() {
		<-ctx.Done()
		delete(user, id)
	}()
	user[id] = events
	return events, nil
}
