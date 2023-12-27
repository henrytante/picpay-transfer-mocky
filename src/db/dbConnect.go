package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:senha123@tcp(localhost:3306)/picpay")
	if err != nil{
		return nil, err
	}
	ping := db.Ping()
	if ping != nil{
		log.Fatal(ping)
	}
	return db, nil
}