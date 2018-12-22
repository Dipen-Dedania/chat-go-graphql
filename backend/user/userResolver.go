package user

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend/dal"
	"github.com/aneri/chat-go-graphql/backend/graph"
)

type Resolver struct{}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Subscription() graph.SubscriptionResolver {
	return &subscriptionResolver{r}
}

var ctxt context.Context

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var users []User
	var user User

	row, err := crConn.Db.Query("SELECT id, name FROM chatApp.userdata")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	defer row.Close()
	return users, nil
}
