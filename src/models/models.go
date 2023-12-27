package models

type UserRegister struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Document  string `json:"document"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Rpassword string `json:"rpassword"`
	Merchant  bool   `json:"merchant"`
}

type TransferType struct {
	Value float64 `json:"value"`
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
}

type Mocky struct {
	Approved string `json:"message"`
}