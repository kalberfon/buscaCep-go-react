package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kalberfon/buscaCep-go-react/api/declarations"
)

type ViaCepApiResponse struct {
	CEP         string `json:"cep"`
	LOGRADOURO  string `json:"logradouro"`
	COMPLEMENTO string `json:"complemento"`
	BAIRRO      string `json:"bairro"`
}

func ApiGetDataCep(cep string) (response declarations.DataCep) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	timeToStart := time.Now()
	resp, error := http.Get(url)
	duration := time.Since(timeToStart)
	if error != nil {
		return response
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	responseViaCep := ViaCepApiResponse{}
	json.Unmarshal([]byte(body), &responseViaCep)

	retorno := response.SetByResponse(
		responseViaCep.CEP,
		responseViaCep.LOGRADOURO,
		responseViaCep.COMPLEMENTO,
		responseViaCep.BAIRRO,
		float64(duration),
	)
	return retorno
}
