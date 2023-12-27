package main

import (
	"picpay/src/db"
	"picpay/src/router"
)

func init()  {
	db.InitDB()
}

func main(){
	router.Router()
}