package main

import (
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
func main() {

	mux := http.NewServeMux()

	log.Println("The server is running on port 3001")

	mux.HandleFunc("/", rootHandler)

	srv := http.Server{
		Addr:    ":" + "3001",
		Handler: mux,
	}
	log.Fatal(srv.ListenAndServe())

}
