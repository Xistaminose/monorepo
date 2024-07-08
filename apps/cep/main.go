package main

import (
	"encoding/json"
	"log"
	"monorepo/apps/cep/services"
	"net/http"
)

func main() {
	http.HandleFunc("/cep", getAddressHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAddressHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["zipcode"]
	if !ok || len(keys[0]) < 1 {
		respondWithError(w, "Missing zipcode parameter")
		return
	}
	zipcode := keys[0]

	address, err := services.FetchAddress(zipcode)
	if err != nil {
		respondWithError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

func respondWithError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(services.AddressResponse{
		Erro:     true,
		Mensagem: message,
	})
}
