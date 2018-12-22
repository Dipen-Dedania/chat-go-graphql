package chatConversation

import (
	"context"
	"sync"

	"github.com/aneri/chat-go-graphql/backend/dal"
)

// Mutation start

var chat map[string]chan ChatConversation
var mu sync.Mutex

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) PostMessage(ctx context.Context, sender_name string, receiver_name string, message string) (ChatConversation, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	chats := ChatConversation{
		SenderName:   sender_name,
		ReceiverName: receiver_name,
		Message:      message,
	}
	if _, err := crConn.Db.Exec("INSERT INTO chatApp.chatconversation (sender_name, receiver_name, message, createdat) VALUES ($1,$2,$3,NOW())", chats.SenderName, chats.ReceiverName, chats.Message); err != nil {
		return ChatConversation{}, err
	}
	// Observe new chat conversation
	for _, observe := range chat {
		observe <- chats
	}
	return chats, nil
}

// Mutation end

// Subscription start

type subscriptionResolver struct{ *Resolver }

func (r *subscriptionResolver) MessagePost(ctx context.Context) (<-chan ChatConversation, error) {
	id := randString(8)

	events := make(chan ChatConversation, 1)

	go func() {
		<-ctx.Done()
		delete(chat, id)
	}()
	chat[id] = events
	return events, nil
}

// Subscription end
