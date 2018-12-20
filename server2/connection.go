package server2

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

type DbConnection struct {
	Db *sql.DB
}

var once sync.Once
var instance *DbConnection

func Connect() (*DbConnection, error) {
	connectionString := "postgresql://root@localhost:26257/training?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	once.Do(func() {
		if err != nil {
			log.Fatal("Error while initializing database", err)
		}
		fmt.Println("Database successfully created")
		instance = &DbConnection{
			Db: db,
		}
	})
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS userchat (id SERIAL PRIMARY KEY,name string NOT NULL,email string,contcat string)"); err != nil {
		log.Fatal("Error while creating table")
	}
	fmt.Println("User table created successfully")
	return instance, err
}
