package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := []byte("Hello world")
		w.Write(msg)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
