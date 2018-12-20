package server2

import (
	"context"
	"fmt"
	"log"
)

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

func (r *mutationResolver) Userjoin(ctx context.Context, name string, email string, contcat string) (*User, error) {
	crConn := ctxt.Value("crConn").(*DbConnection)
	defer crConn.Db.Close()
	users := User{
		Name:    name,
		Email:   email,
		Contact: contcat,
	}
	if _, err := crConn.Db.Exec("INSERT INTO training.userchat (name,email,contcat) VALUES ($1,$2,$3)", users.Name, users.Email, users.Contact); err != nil {
		log.Fatal("Error while inserting data", err)
	}
	fmt.Println("Data inserted sucessfully")
	return users, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) UserJoined(ctx context.Context) (<-chan *User, error) {
	panic("not implemented")
}
