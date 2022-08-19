package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseError(w http.ResponseWriter, msg string, status int) {
	err := json.NewEncoder(w).Encode(map[string]string{"error": msg})
	if err != nil {
		log.Println("error writing json", err)
		return
	}
	w.WriteHeader(status)
}
