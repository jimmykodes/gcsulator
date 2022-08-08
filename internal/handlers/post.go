package handlers

import (
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// todo: implement upload logic
	responseError(w, "not implemented", http.StatusInternalServerError)
}
