package resolver

import (
	"context"
	"log"

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
			log.Print("Error in scanning data of users", err)
		}
		arrusers = append(arrusers, user)
	}
	defer rows.Close()
	return arrusers, nil
}
func (r *mutationResolver) JoinUser(ctx context.Context, name string) (model.User, error) {
	crConn := ctxt.Value("crConn").(*dal.DbConnection)
	users := model.User{
		Name: name,
	}
	row, err := crConn.Db.Query("SELECT name FROM userdata WHERE name=$1", name)
	if err != nil {
		log.Print("Error while scanning name")
	}
	for row.Next() {
		err := row.Scan(&users.Name)
		if err != nil {
			log.Print(err)
		}
	}
	if users.Name != "" {
		_, err := crConn.Db.Exec("INSERT INTO userdata (name, createdat) VALUES ($1, NOW())", name)
		if err != nil {
			log.Print("Error while inserting data", err)
		}
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
