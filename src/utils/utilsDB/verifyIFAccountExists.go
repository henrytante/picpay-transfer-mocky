package utilsdb

import (
	"picpay/src/db"
)

func VerifyIfAccountExists(payer, payee int) (bool, error) {
    db, err := db.ConnectDB()
    if err != nil {
        return false, err
    }
    defer db.Close()
    var countPayer, countPayee int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", payer).Scan(&countPayer)
    if err != nil {
        return false, err
    }
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", payee).Scan(&countPayee)
    if err != nil {
        return false, err
    }
    return countPayer > 0 && countPayee > 0, nil
}
