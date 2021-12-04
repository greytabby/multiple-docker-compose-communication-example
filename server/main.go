package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request received from ", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
		return
	})
	log.Print("Start server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
