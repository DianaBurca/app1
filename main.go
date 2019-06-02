package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		output := map[string]string{
			"message": "I am the second app and I work fine",
			"status":  "200",
		}
		json.NewEncoder(w).Encode(output)
	})

	http.HandleFunc("/.well-known/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/.well-known/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
