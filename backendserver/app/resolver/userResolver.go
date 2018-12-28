package resolver

import (
	"context"

	"github.com/aneri/chat-go-graphql/backendserver/app/dal"
	"github.com/aneri/chat-go-graphql/backendserver/app/model"
)

type userResolver struct{ *Resolver }

var addUserChannel map[string]chan model.User

func init() {
	addUserChannel = map[string]chan model.User{}
}

func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
	var arrusers []model.User
	var user model.User
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	rows, err := crConn.Db.Query("SELECT id,name,createdat FROM userdata")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		arrusers = append(arrusers, user)
	}
	defer rows.Close()
	return arrusers, nil
}
func (r *mutationResolver) JoinUser(ctx context.Context, name string) (model.User, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	_, err := crConn.Db.Exec("INSERT INTO userdata (name, createdat) VALUES ($1, NOW())", name)
	if err != nil {
		panic(err.Error())
	}
	users := model.User{
		Name: name,
	}
	//add new user in observer
	for _, observer := range addUserChannel {
		observer <- users
	}
	return users, nil
}
func (r *subscriptionResolver) Userjoined(ctx context.Context) (<-chan model.User, error) {
	id := randString(8)
	userEvent := make(chan model.User, 1)
	go func() {
		<-ctx.Done()
		delete(addUserChannel, id)
	}()
	addUserChannel[id] = userEvent
	return userEvent, nil
}
