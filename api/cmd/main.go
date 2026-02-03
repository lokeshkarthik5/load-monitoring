package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	log.Println("The server is running on port 3001")

	srv := http.Server{
		Addr:    ":" + "3001",
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())

}
