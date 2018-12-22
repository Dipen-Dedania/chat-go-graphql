package chatConversation

import (
	"time"
)

type ChatConversation struct {
	ID           string    `json:"id"`
	SenderName   string    `json:"sender_name"`
	ReceiverName string    `json:"receiver_name"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"createdAt"`
}