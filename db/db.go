package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	conexao := "user=postgres dbname=GW_store password=Grasieli11 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
