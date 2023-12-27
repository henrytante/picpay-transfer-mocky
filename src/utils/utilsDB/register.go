package utilsdb

import (
	"picpay/src/db"
	"picpay/src/models"
)


func RegisterUser(data *models.UserRegister) error {
	db, _ := db.ConnectDB()
	defer db.Close()
	prepare, err := db.Prepare("INSERT INTO users (nome, sobrenome, documento, email, senha, merchant) VALUES (?,?,?,?,?,?)")
	if err != nil{
		return err
	}
	defer prepare.Close()
	insert, err := prepare.Exec(data.FirstName, data.LastName, data.Document, data.Email, data.Password,  data.Merchant)
	if err != nil{
		return err
	}
	userID, err := insert.LastInsertId()
	if err != nil{
		return err
	}
	prepare, err = db.Prepare("INSERT INTO saldos (id) VALUES (?)")
	if err != nil{
		return err
	}
	defer prepare.Close()
	if _, err = prepare.Exec(userID); err != nil{
		return err
	}
	return nil
}