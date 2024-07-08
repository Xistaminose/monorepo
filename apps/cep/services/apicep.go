package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiCEPFetcher struct{}

type ApiCEPResponse struct {
	CEP        string `json:"code"`
	Logradouro string `json:"address"`
	Bairro     string `json:"district"`
	Localidade string `json:"city"`
	UF         string `json:"state"`
}

func (a ApiCEPFetcher) Fetch(zipcode string) (*AddressResponse, error) {
	url := fmt.Sprintf("https://ws.apicep.com/cep/%s.json", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ApiCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if result.Logradouro == "" {
		return nil, fmt.Errorf("no address found for this zipcode ApiCep")
	}

	return &AddressResponse{
		CEP:        result.CEP,
		Logradouro: result.Logradouro,
		Bairro:     result.Bairro,
		Localidade: result.Localidade,
		UF:         result.UF,
		Erro:       false,
		Source:     "ApiCEP",
	}, nil
}
