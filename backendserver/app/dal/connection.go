package dal

import (
	"database/sql"
	"fmt"
	"log"

	"sync"

	_ "github.com/lib/pq"
)

// DbConnection model
type DbConnection struct {
	Db *sql.DB
}

var once sync.Once
var instance *DbConnection

// DbConnect for database connection
func DbConnect() (*DbConnection, error) {

	once.Do(func() {
		fmt.Println("starting server")
		configuration, err := LoadConfiguration("app/dal/config.json")
		if err != nil {
			log.Fatal("Error while starting server", err)
		}
		connectionString := fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", configuration.Cockroach.User, configuration.Cockroach.Host, configuration.Cockroach.Port, configuration.Cockroach.DbName)
		db, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatal("error while initializing database", err)
		}
		fmt.Println("Database successfulyy initialized")
		instance = &DbConnection{
			Db: db,
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS userdata (id SERIAL PRIMARY KEY,name string NOT NULL UNIQUE,createdat timestamptz)"); err != nil {
			log.Fatal("error while creating user table", err)
		}
		if _, err := db.Exec("CREATE TABLE IF NOT EXISTS chatconversation (id SERIAL PRIMARY KEY,sender_name string NOT NULL REFERENCES userdata (name),receiver_name string NOT NULL REFERENCES userdata (name),message string NOT NULL,createdat timestamptz)"); err != nil {
			log.Fatal("error while creating chat table", err)
		}
	})
	return instance, nil
}
