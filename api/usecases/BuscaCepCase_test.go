package usecases

import (
	"reflect"
	"testing"

	"github.com/kalberfon/buscaCep-go-react/api/declarations"
)

func TestBuscaCepCase(t *testing.T) {
	websites := []string{
		"57035226",
	}

	resultado := BuscaCepCase(websites)

	esperado := map[string]declarations.DataCep{
		"57035226": {CEP: "57035226"},
	}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado inesperado")
	}
}
