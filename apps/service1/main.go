package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a http server that return "Hello from service1"

	fmt.Println("Hello from service1")
	fmt.Println("Hello2 from common")

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello from service1"))
		}),
	}

	server.ListenAndServe()

}
