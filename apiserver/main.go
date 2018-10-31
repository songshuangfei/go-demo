package main

import (
	"go-demo/apiserver/handler"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/login", handler.Login)
	router.GET("/user/:uid", handler.User)

	router.GET("/api/:apiType", handler.Api)
	router.NotFound = http.HandlerFunc(handler.NotFound)
	log.Fatal(http.ListenAndServe(":8080", router))
}
