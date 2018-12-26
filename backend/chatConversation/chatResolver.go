package chatConversation

import (
	"context"

	"github.com/aneri/chat-go-graphql/backend/dal"
	"github.com/aneri/chat-go-graphql/backend/graph"
)

var ctxt context.Context

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

type queryResolver struct{ *Resolver }

func (r *queryResolver) Chat(ctx context.Context, sender_name string, receiver_name string) ([]ChatConversation, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var chats []ChatConversation
	var chat ChatConversation
	row, err := crConn.Db.Query("SELECT sender_name, receiver_name, message, createdat FROM chatApp.chatconversation WHERE sender_name IN ($1,$2) and receiver_name IN ($1,$2) order by chatconversation.id", chat.SenderName, chat.ReceiverName)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&chat.SenderName, &chat.ReceiverName, &chat.Message, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	defer row.Close()
	return chats, nil
}
