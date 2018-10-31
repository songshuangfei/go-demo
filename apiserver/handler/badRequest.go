package handler

import (
	"io"
	"net/http"
)

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
	io.WriteString(w, "请求错误")
}
