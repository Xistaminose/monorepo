package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenCEPFetcher struct{}

type OpenCEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func (o OpenCEPFetcher) Fetch(zipcode string) (*AddressResponse, error) {
	url := fmt.Sprintf("https://opencep.com/v1/%s", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result OpenCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Logradouro == "" {
		return nil, fmt.Errorf("no address found for this zipcode OpenCEP")
	}

	return &AddressResponse{
		CEP:        result.CEP,
		Logradouro: result.Logradouro,
		Bairro:     result.Bairro,
		Localidade: result.Localidade,
		UF:         result.UF,
		Erro:       false,
		Source:     "OpenCEP",
	}, nil
}
