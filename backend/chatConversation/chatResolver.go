package chatConversation

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend"
)

// Query start
func (r *backend.Resolver) Query() backend.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *backend.Resolver }

func (r *queryResolver) Chat(ctx context.Context, sender_name string, receiver_name string) ([]ChatConversation, error) {
	panic("not implemented")
}

// Query end
