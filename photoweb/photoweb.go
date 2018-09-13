package main

import (
	"go-demo/photoweb/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", handler.UploadHandler)
	http.HandleFunc("/view", handler.ViewHandler)
	http.HandleFunc("/", handler.ListHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}
}
