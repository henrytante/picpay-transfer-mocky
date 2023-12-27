package utilsdb

import (
	"errors"
	"picpay/src/db"
)

func Transfer(payer, payee int, value float64) error {
	dbConn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	accountsExist, err := VerifyIfAccountExists(payer, payee)
	if err != nil {
		return err
	}
	if !accountsExist {
		return errors.New("Uma das contas dessa transferência não existe")
	}
	isMerchant, err := GetUserType(payer)
	if err != nil {
		return err
	}
	if isMerchant {
		return errors.New("Vendedores não podem fazer transferências, apenas receber")
	}
	_, err = dbConn.Exec("UPDATE saldos SET saldo = saldo + ? WHERE id = ?", value, payee)
	if err != nil {
		return err
	}
	_, err = dbConn.Exec("UPDATE saldos SET saldo = saldo - ? WHERE id = ?", value, payer)
	if err != nil {
		_, rollbackErr := dbConn.Exec("UPDATE saldos SET saldo = saldo - ? WHERE id = ?", value, payee)
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return nil
}
