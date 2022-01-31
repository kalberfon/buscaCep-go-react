package declarations

type DataCep struct {
	CEP         string `json:"cep"`
	LOGRADOURO  string
	COMPLEMENTO string
	BAIRRO      string
	DURATION    float64
}

func (dc DataCep) SetByResponse(cep string, logradouro string, complemento string, bairro string, duration float64) DataCep {
	dc.BAIRRO = bairro
	dc.LOGRADOURO = logradouro
	dc.COMPLEMENTO = complemento
	dc.CEP = cep
	dc.DURATION = duration
	return dc
}
func (dc DataCep) ToString() string {
	return "cep: " + dc.CEP + " lgoradouro:" + dc.LOGRADOURO + " complemento: " + dc.COMPLEMENTO + " bairro: " + dc.BAIRRO
}
