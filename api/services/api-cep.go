package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kalberfon/buscaCep-go-react/api/declarations"
)

type ApiCepResponse struct {
	STATUS     string `json:"status"`
	CEP        string `json:"code"`
	CITY       string `json:"city"`
	LOGRADOURO string `json:"address"`
	BAIRRO     string `json:"district"`
}

func ApiCep(cep string) (retorno declarations.DataCep) {
	url := "https://ws.apicep.com/cep/" + cep + ".json"
	startTiem := time.Now()

	reponse, error := http.Get(url)
	duration := time.Since(startTiem)
	if error != nil {
		return retorno
	}
	defer reponse.Body.Close()

	body, _ := ioutil.ReadAll(reponse.Body)
	apiResponse := ApiCepResponse{}
	json.Unmarshal([]byte(body), &apiResponse)

	return retorno.SetByResponse(apiResponse.CEP, apiResponse.LOGRADOURO, "", apiResponse.BAIRRO, float64(duration))
}
