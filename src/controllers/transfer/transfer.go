package transfer

import (
	"encoding/json"
	"net/http"
	"picpay/src/models"
	"picpay/src/utils/mocky"
	"picpay/src/utils/send"
	utilsdb "picpay/src/utils/utilsDB"
)

func Transfer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		send.JSON(w, http.StatusMethodNotAllowed, "Metodo não permitido")
		return
	}
	var transferData models.TransferType
	err := json.NewDecoder(r.Body).Decode(&transferData)
	if err != nil {
		send.JSON(w, http.StatusUnprocessableEntity, "Erro ao decodificar os dados")
		return
	}
	approved, err := mocky.Approve()
	if err != nil {
		send.JSON(w, http.StatusBadGateway, "Erro na conexão com o mocky ou transferencia não aprovada")
		return
	}
	if approved {
		err = utilsdb.Transfer(transferData.Payer, transferData.Payee, transferData.Value)
		if err != nil {
			send.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}
		send.JSON(w, http.StatusOK, "Tranferencia realizada com sucesso")
	}
}
