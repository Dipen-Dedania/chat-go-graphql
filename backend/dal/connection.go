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
	configuration, err := config.LoadConfiguration("config.json")
	if err != nil {
		return nil, err
	}
	connectionString := fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", configuration.Cockroach.User, configuration.Cockroach.Host, configuration.Cockroach.Port, configuration.Cockroach.DbName)
	db, err := sql.Open("postgres", connectionString)
	once.Do(func() {
		if err != nil {
			log.Fatal("error while initializing database")
		}
		fmt.Println("Database successfulyy created")
		instance = &DbConnection{
			Db: db,
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS chatApp.userdata (id SERIAL PRIMARY KEY,name string NOT NULL,email string,contactno string,password string)"); err != nil {
			log.Fatal("error while creating user table", err)
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS chatApp.chat (id SERIAL PRIMARY KEY, sender_id integer NOT NULL REFERENCES chatApp.userdata (id), receiver_id integer NOT NULL chatApp.userdata (id), message string NOT NULL,createdAt timestamptz)"); err != nil {
			log.Fatal("error while creating chat table", err)
		}
		fmt.Println("Tables created successfully")
	})
	return instance, err
}
