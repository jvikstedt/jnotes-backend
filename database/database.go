package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Required driver
)

// DB database connection
var DB *sqlx.DB

// Setup database connection
func Setup() {
	var err error
	DB, err = sqlx.Connect("postgres", "postgres://jnotes:jnotes@localhost/jnotes_development")
	if err != nil {
		panic(err)
	}
}
