package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDBContext(connStr string) *sql.DB {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to db")
	return db
}
