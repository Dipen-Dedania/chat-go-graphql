package backend

import (
	"context"
)

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

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input UpdateUser) (User, error) {
	panic("not implemented")
}
func (r *mutationResolver) PostMessage(ctx context.Context, senderid int, receiverid int, message string) (*Chat, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) Chats(ctx context.Context, client1name string, client2name string) (Chat, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) MessagePosted(ctx context.Context) (<-chan Chat, error) {
	panic("not implemented")
}
