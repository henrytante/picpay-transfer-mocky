package db

import "log"

func InitDB() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			nome VARCHAR(30) NOT NULL,
			sobrenome VARCHAR(150) NOT NULL,
			documento VARCHAR(12) NOT NULL UNIQUE,
			email VARCHAR(150) NOT NULL UNIQUE,
			senha VARCHAR(150) NOT NULL,
			merchant boolean default false
		)
	`

	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}

	createTableQuerySaldo := `
		CREATE TABLE IF NOT EXISTS saldos (
			id INT PRIMARY KEY,
			saldo FLOAT DEFAULT 0,
			FOREIGN KEY (id) REFERENCES users(id)
		)
	`
	if _, err := db.Exec(createTableQuerySaldo); err != nil {
		log.Fatal(err)
	}
}
