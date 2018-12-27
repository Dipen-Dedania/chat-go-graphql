package backend

import (
	"context"
	"log"

	"github.com/aneri/chat-go-graphql/backend/dal"
)

var chat map[string]chan Chatconversation

func (r *mutationResolver) PostMessage(ctx context.Context, sender_name string, receiver_name string, message string) (Chatconversation, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	chats := Chatconversation{
		SenderName:   sender_name,
		ReceiverName: receiver_name,
		Message:      message,
	}
	if _, err := crConn.Db.Exec("INSERT INTO chatconversation (sender_name, receiver_name, message, createdat) VALUES ($1,$2,$3,NOW())", chats.SenderName, chats.ReceiverName, chats.Message); err != nil {
		log.Fatal("Error while inserting data in chat table", err)
		return Chatconversation{}, err
	}
	// Observer new chat
	for _, observe := range chat {
		observe <- chats
	}
	return chats, nil
}

func (r *queryResolver) Chats(ctx context.Context, sender_name string, receiver_name string) ([]Chatconversation, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	var chats []Chatconversation
	var chat Chatconversation
	row, err := crConn.Db.Query("SELECT id,sender_name,receiver_name,message,createdat FROM chatconversation WHERE sender_name IN ($1,$2) and receiver_name IN ($1,$2) ORDER BY createdat", sender_name, receiver_name)
	if err != nil {
		log.Fatal("Error while retrieving chat data")
		return nil, err
	}
	for row.Next() {
		if err := row.Scan(&chat.ID, &chat.SenderName, &chat.ReceiverName, &chat.Message, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, id string) (<-chan Chatconversation, error) {
	// create new channel for request
	chatevent := make(chan Chatconversation, 1)
	chat[id] = chatevent

	// Delete channel when done
	go func() {
		<-ctx.Done()
		delete(chat, id)
	}()
	return chatevent, nil
}
