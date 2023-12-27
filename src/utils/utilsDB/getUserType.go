package utilsdb

import (
	"database/sql"
	"errors"
	"picpay/src/db"
)

func GetUserType(id int) (bool, error){
	db, _ := db.ConnectDB()
	defer db.Close()
	var merchant bool
	err := db.QueryRow("SELECT merchant FROM users WHERE id = ?", id).Scan(&merchant)
	if err != nil{
		if err == sql.ErrNoRows{
			return false, errors.New("Usuario não encontrado")
		}
		return false, err
	}
	return merchant, nil
}
