package main

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/hello-world/user"
)

func main() {
	http.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": Hello(),
			"user":    user.GetName(),
		})
	})

	log.Println("Server running at http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
