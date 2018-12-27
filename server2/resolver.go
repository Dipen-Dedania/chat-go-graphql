package server2

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

var user map[string]chan User
var mu sync.Mutex
var ctxt context.Context

type Resolver struct{}

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
func (r *mutationResolver) Userjoin(ctx context.Context, name string, email *string, contact *string) (User, error) {
	// Database connection
	crConn := ctxt.Value("crConn").(*DbConnection)
	// create User
	users := User{
		Name: name,
	}
	if _, err := crConn.Db.Exec("INSERT INTO training.userchat (name, email, contact) VALUES ($1,$2,$3)", users.Name, "", ""); err != nil {
		return User{}, err
	}
	defer crConn.Db.Close()
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

	row, err := crConn.Db.Query("SELECT id, name, email ,contact FROM training.userchat")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Contact); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	defer row.Close()
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
