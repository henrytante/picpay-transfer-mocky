package utilsdb

import (
	"database/sql"
	"errors"
	"picpay/src/db"
)

func GetUserSaldo(id int) (float64, error) {
	db, _ := db.ConnectDB()
	defer db.Close()
	var saldo float64
	err := db.QueryRow("SELECT saldo FROM saldos WHERE id = ?", id).Scan(saldo)
	if err != nil{
		if err == sql.ErrNoRows{
			return 0, errors.New("Usuario não encontrado")
		}
		return 0, err
	}
	return saldo, nil
}