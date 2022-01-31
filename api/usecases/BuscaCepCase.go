package usecases

import (
	"github.com/kalberfon/buscaCep-go-react/api/declarations"
	"github.com/kalberfon/buscaCep-go-react/api/services"
)

type ChannelResponse struct {
	ENDPOINT string
	DATA     declarations.DataCep
}

func BuscaCepCase(ceps []string) map[string][]declarations.DataCep {
	resposta := make(map[string][]declarations.DataCep)

	canalResposta := make(chan ChannelResponse)
	canalRespostaApicep := make(chan ChannelResponse)
	for _, cep := range ceps {
		go func(cep string) { // goroutine for viacep consulting
			canalResposta <- ChannelResponse{ENDPOINT: "viacep", DATA: services.ApiGetDataCep(cep)}
		}(cep)
		go func(cep string) {
			canalRespostaApicep <- ChannelResponse{ENDPOINT: "apicep", DATA: services.ApiCep(cep)}
		}(cep)
	}

	for i := 0; i < len(ceps); i++ {
		resp := <-canalResposta
		resposta[resp.ENDPOINT] = append(resposta[resp.ENDPOINT], resp.DATA)
	}
	for i := 0; i < len(ceps); i++ {
		resp := <-canalRespostaApicep
		resposta[resp.ENDPOINT] = append(resposta[resp.ENDPOINT], resp.DATA)
	}

	return resposta
}
