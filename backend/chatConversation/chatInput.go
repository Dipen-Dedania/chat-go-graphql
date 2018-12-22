package chatConversation

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend"
)

// Mutation start
func (r *backend.Resolver) Mutation() backend.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *backend.Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, sender_name string, receiver_name string, message string) (ChatConversation, error) {
	panic("not implemented")
}

// Mutation end

// Subscription start
func (r *backend.Resolver) Subscription() backend.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type subscriptionResolver struct{ *backend.Resolver }

func (r *subscriptionResolver) MessagePost(ctx context.Context) (<-chan ChatConversation, error) {
	panic("not implemented")
}

// Subscription end
