package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jimmykodes/prmrtr"
)

func Get(w http.ResponseWriter, r *http.Request) {
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

	f, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()

	if _, err := io.Copy(w, f); err != nil {
		responseError(w, "error sending file", http.StatusInternalServerError)
		return
	}
}
