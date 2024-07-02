package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDbConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=example host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
