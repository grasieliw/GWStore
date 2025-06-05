package db

import (
	"database/sql"

	"os"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
