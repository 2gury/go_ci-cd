package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mx := mux.NewRouter()
	mx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	}).Methods(http.MethodGet)

	
	log.Println("Start server at :8080 port")
	http.ListenAndServe(":8080", mx)
}
