package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(map[string]string{"error": msg})
	if err != nil {
		log.Println("error writing json", err)
	}
}
