package router

import (
	"net/http"
	"picpay/src/controllers/register"
	"picpay/src/controllers/transfer"
)


func Router()  {
	http.HandleFunc("/auth/register", register.RegisterUser)
	http.HandleFunc("/transfer", transfer.Transfer)

	http.ListenAndServe(":8080", nil)
}