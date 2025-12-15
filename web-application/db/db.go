package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBanco() *sql.DB {
	conn := "user=postgres dbname=loja password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	return db
}
