package db

import (
	"database/sql"

	"log"
	"os"

	_ "github.com/lib/pq"
)

var stringConexao = os.Getenv("DATABASE_URL")

var DB *sql.DB

func ConectaBancoDados() *sql.DB {

	db, err := sql.Open("postgres", stringConexao)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	return db

}
