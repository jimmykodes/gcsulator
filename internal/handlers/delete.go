package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/jimmykodes/prmrtr"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := prmrtr.Vars(r)
	bucket, ok := vars.String("bucket")
	if !ok {
		responseError(w, "missing bucket", http.StatusBadRequest)
		return
	}
	object, ok := vars.String("object")
	if !ok {
		responseError(w, "missing object", http.StatusBadRequest)
		return
	}
	path := filepath.Join("/var/gcsulator", bucket, object)
	if err := os.Remove(path); err != nil {
		responseError(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
