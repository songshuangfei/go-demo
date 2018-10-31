package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//User handler user
func User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}
