package mocky

import (
	"encoding/json"
	
	"net/http"
	"picpay/src/models"
)

func Approve() (bool, error) {
	url := "https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return false, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return false, err
	}
	defer resp.Body.Close()
	var response models.Mocky
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil{
		return false, err
	}
	switch response.Approved{
	case "Autorizado":
		return true, nil
	default:
		return false, err
	}
}