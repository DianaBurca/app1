package main

import (
	"encoding/json"
	"net/http"
)

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

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

	http.HandleFunc("/compute", func(w http.ResponseWriter, r *http.Request) {
		result := map[uint64]uint64{}

		for i := uint64(0); i < 10; i++ {
			result[i] = Factorial(uint64(i))
		}

		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
