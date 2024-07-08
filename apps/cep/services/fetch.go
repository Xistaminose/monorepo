package services

import (
	"fmt"
	"log"
)

type AddressResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Erro       bool   `json:"erro"`
	Mensagem   string `json:"mensagem,omitempty"`
	Source     string `json:"source,omitempty"`
}

type Fetcher interface {
	Fetch(zipcode string) (*AddressResponse, error)
}

func FetchAddress(zipcode string) (*AddressResponse, error) {
	fetchers := []Fetcher{
		ViaCEPFetcher{},
		ApiCEPFetcher{},
		OpenCEPFetcher{},
	}

	results := make(chan *AddressResponse, len(fetchers))
	errors := make(chan error, len(fetchers))

	for _, fetcher := range fetchers {
		go func(f Fetcher) {
			address, err := f.Fetch(zipcode)
			if err != nil {
				errors <- err
				return
			}
			results <- address
		}(fetcher)
	}

	for range fetchers {
		select {
		case res := <-results:
			return res, nil
		case err := <-errors:
			log.Println("Error:", err)
		}
	}

	return nil, fmt.Errorf("no valid address found")
}
