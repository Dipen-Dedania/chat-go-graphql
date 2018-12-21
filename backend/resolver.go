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

func (r *mutationResolver) JoinUser(ctx context.Context, name string) (User, error) {
	panic("not implemented")
}errerrerr
func (r *mutationResolver) PostMessage(ctx context.Context, sender_name string, receiver_name string, message string) (ChatConversation, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) Chat(ctx context.Context, sender_name string, receiver_name string) ([]ChatConversation, error) {
	panic("not implemented")
}

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) UserJoined(ctx context.Context) (<-chan User, error) {
	panic("not implemented")
}
func (r *subscriptionResolver) MessagePost(ctx context.Context) (<-chan ChatConversation, error) {
	panic("not implemented")
}
