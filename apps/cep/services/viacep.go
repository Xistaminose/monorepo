package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ViaCEPFetcher struct{}

type ViaCEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       bool   `json:"erro"`
}

func (v ViaCEPFetcher) Fetch(zipcode string) (*AddressResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Erro {
		return nil, fmt.Errorf("no address found for this zipcode ViaCEP")
	}

	return &AddressResponse{
		CEP:        result.CEP,
		Logradouro: result.Logradouro,
		Bairro:     result.Bairro,
		Localidade: result.Localidade,
		UF:         result.UF,
		Erro:       false,
		Source:     "ViaCEP",
	}, nil
}
