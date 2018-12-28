package resolver

import (
	"context"
	"log"

	"github.com/aneri/chat-go-graphql/backendserver/app/dal"
	"github.com/aneri/chat-go-graphql/backendserver/app/model"
)

type chatResolver struct{ *Resolver }

var addChatResolver map[string]chan model.Chatconversation

func init() {
	addChatResolver = map[string]chan model.Chatconversation{}
}

func (r *queryResolver) Chats(ctx context.Context, sender_name string, receiver_name string) ([]model.Chatconversation, error) {
	var arrChats []model.Chatconversation
	var chat model.Chatconversation
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	rows, err := crConn.Db.Query("SELECT id,sender_name,receiver_name,message,createdat FROM chatconversation WHERE sender_name IN ($1,$2) and receiver_name IN ($1,$2) ORDER BY createdat", sender_name, receiver_name)
	if err != nil {
		log.Fatal("Error while retrieving chat data", err)
	}
	for rows.Next() {
		err := rows.Scan(&chat.ID, &chat.SenderName, &chat.ReceiverName, &chat.Message, &chat.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		arrChats = append(arrChats, chat)
	}
	defer rows.Close()
	return arrChats, nil
}

func (r *mutationResolver) PostMessage(ctx context.Context, sender_name string, receiver_name string, message string) (model.Chatconversation, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	if _, err := crConn.Db.Exec("INSERT INTO chatconversation (sender_name, receiver_name, message, createdat) VALUES ($1,$2,$3,NOW())", sender_name, receiver_name, message); err != nil {
		log.Fatal("Error while inserting data in chat table", err)
	}
	chats := model.Chatconversation{
		ReceiverName: receiver_name,
		SenderName:   sender_name,
		Message:      message,
	}
	// add new chat observe
	for _, observer := range addChatResolver {
		observer <- chats
	}
	return chats, nil
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, id string) (<-chan model.Chatconversation, error) {
	// create new channel for request
	chatevent := make(chan model.Chatconversation, 1)
	addChatResolver[id] = chatevent
	go func() {
		<-ctx.Done()
		delete(addChatResolver, id)
	}()
	return chatevent, nil
}
