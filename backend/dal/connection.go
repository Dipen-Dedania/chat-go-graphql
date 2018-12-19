package dal

import (
	"database/sql"
	"sync"
)

type DbConnection struct {
	Db *sql.DB
}

var once sync.Once
var instance *DbConnection
