package handler

import (
	"io"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	notFoundHandler(w, r)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	io.WriteString(w, r.URL.String()+" is not found")
}
