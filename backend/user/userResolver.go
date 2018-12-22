package user

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend"
)

func (r *backend.Resolver) Query() backend.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *backend.Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
