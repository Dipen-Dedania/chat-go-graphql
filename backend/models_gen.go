// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package backend

import (
	"time"
)

type Chat struct {
	ID         string    `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"createdAt"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Contact  string `json:"contact"`
	Password string `json:"password"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
