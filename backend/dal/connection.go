package dal

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/aneri/chat-go-graphql/backend/config"
)

// DbConnection model
type DbConnection struct {
	Db *sql.DB
}

var once sync.Once
var instance *DbConnection

// DbConnect for database connection
func DbConnect() (*DbConnection, error) {
	fmt.Println("starting server")
	config, err := config.LoadConfiguration("config.json")
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s host=%s port=%s dbname=%s sslmode=disable", config.Cockroach.User, config.Cockroach.Host, config.Cockroach.Port, config.Cockroach.DbName))
	once.Do(func() {
		fmt.Printf("user=%s host=%s port=%s dbname=%s sslmode=disable", config.Cockroach.User, config.Cockroach.Host, config.Cockroach.Port, config.Cockroach.DbName)
		if err != nil {
			log.Fatal("error while initializing database")
		}
		fmt.Println("Database successfulyy created")
		instance = &DbConnection{
			Db: db,
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS userdata (id SERIAL PRIMARY KEY,name string NOT NULL,email string,contactno string,password string)"); err != nil {
			return nil, err
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS chat (id SERIAL PRIMARY KEY, sender_id integer NOT NULL REFERENCES chatApp.userdata (id), receiver_id integer NOT NULL chatApp.userdata (id), message string NOT NULL,createdAt timestamptz)"); err != nil {
			return nil, err
		}
	})
	return instance, err
}
