package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	loginHandler(w, r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
}
