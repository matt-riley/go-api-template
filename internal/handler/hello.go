package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"}); err != nil {
		log.Printf("write hello response: %v", err)
	}
}
