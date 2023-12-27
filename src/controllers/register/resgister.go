package register

import (
	"encoding/json"
	"net/http"
	"picpay/src/models"
	"picpay/src/utils/send"
	utilsdb "picpay/src/utils/utilsDB"
	"github.com/go-sql-driver/mysql"
)



func RegisterUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost{
		send.JSON(w, http.StatusMethodNotAllowed, "Metodo não permitido")
		return
	}
	var userData models.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil{
		send.JSON(w, http.StatusInternalServerError, "Erro ao decodificar os dados")
		return
	}
	if userData.FirstName == "" || userData.LastName == "" || userData.Document == "" || userData.Email == "" || userData.Password == "" || userData.Rpassword == ""{
		send.JSON(w, http.StatusBadRequest, "Dados em branco")
		return
	}
	if userData.Password != userData.Rpassword{
		send.JSON(w, http.StatusBadRequest, "As senhas não batem")
		return
	}
	err := utilsdb.RegisterUser(&userData)
	if err != nil{
		if SqlError, ok := err.(*mysql.MySQLError); ok {
			if SqlError.Number == 1062{
				send.JSON(w, http.StatusConflict, "Já existe um usuario com esses dados")
				return
			}
		}
		send.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	send.JSON(w, http.StatusOK, "Usuario registrado")
}