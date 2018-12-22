package user

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend"
)

func (r *backend.Resolver) Mutation() backend.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *backend.Resolver }

func (r *mutationResolver) JoinUser(ctx context.Context, name string) (User, error) {
	panic("not implemented")
}
func (r *backend.Resolver) Subscription() backend.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *backend.Resolver }

func (r *subscriptionResolver) UserJoined(ctx context.Context) (<-chan User, error) {
	panic("not implemented")
}
