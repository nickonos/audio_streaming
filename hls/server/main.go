package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", addHeaders(http.FileServer(http.Dir("audio"))))

	fmt.Printf("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
