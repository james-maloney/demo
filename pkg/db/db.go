package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	log.SetFlags(log.Llongfile)
}

// Init is used to setup a DB connection.
func Init(username, password, dbName string) {
	ip := "127.0.0.1:3306"

	var err error
	RW.DB, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		username,
		password,
		ip,
		dbName,
	))
	if err != nil {
		log.Fatal(err)
	}

	// make sure we can access the DB
	if err := RW.Ping(); err != nil {
		log.Fatal(err)
	}

	RW.connected = true
}

// RW read/write connection
var RW = DB{
	DB: &sqlx.DB{},
}

func (db DB) IsConnected() bool {
	return db.connected
}

// DB wraps sqlx.DB so we can override methods if we need to
type DB struct {
	*sqlx.DB
	connected bool
}
