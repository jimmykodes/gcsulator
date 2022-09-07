package handlers

import (
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("route no found", r.Method, r.URL.Path)
	w.WriteHeader(http.StatusNotFound)
}
