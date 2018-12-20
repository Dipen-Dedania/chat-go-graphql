package server2

import (
	"context"
	"log"
	"net/http"
	"sync"
)

type Resolver struct{ *Resolver }

var user map[string]chan User
var mu sync.Mutex
var ctxt context.Context

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }

func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		crConn, err := Connect()
		if err != nil {
			log.Fatal(err)
		}
		ctxt = context.WithValue(request.Context(), "crConn", crConn)
		next.ServeHTTP(writer, request.WithContext(ctxt))
	})
}
func (r *mutationResolver) Userjoin(ctx context.Context, name string, email string, contact string) (User, error) {
	// Database connection
	crConn := ctxt.Value("crConn").(*DbConnection)
	// create User
	users := User{
		Name:    name,
		Email:   email,
		Contact: contact,
	}
	if _, err := crConn.Db.Exec("INSERT INTO training.userchat (name,email,contact) VALUES ($1,$2,$3)", users.Name, users.Name, users.Contact); err != nil {
		return User{}, err
	}
	// Observer new user joined
	for _, observe := range user {
		observe <- users
	}
	// return user data
	return users, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	crConn := ctxt.Value("crConn").(*DbConnection)
	var users []User
	var user User

	row, err := crConn.Db.Query("SELECT (id, name, email ,contact) FROM training.userchat")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&user.ID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	defer row.Close()
	return users, nil
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) UserJoined(ctx context.Context) (<-chan User, error) {

	id := GetId()
	events := make(chan User, 1)
	go func() {
		<-ctx.Done()
		delete(user, id)
	}()
	user[id] = events
	return events, nil
}

func GetId() string {
	crConn := ctxt.Value("crConn").(*DbConnection)
	var user User
	row, err := crConn.Db.Query("SELECT (id) FROM training.chatuser")
	if err != nil {
		log.Fatal("error while fetching id", err)
	}
	row.Scan(&user.ID)
	defer row.Close()
	return user.ID
}
