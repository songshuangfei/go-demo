package handler

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type apiHandler func(http.ResponseWriter, *http.Request)

var apiHandlers map[string]apiHandler

//Api handler ApiHome
func Api(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	apiType := ps.ByName("apiType")
	if handler, ok := apiHandlers[apiType]; ok {
		handler(w, r)
		return
	}
	NotFound(w, r)
}

//handler函数
func carouselHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "carousel")
	io.WriteString(w, r.URL.Query().Encode())
}

func hotAuthors(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hot Authors")
	io.WriteString(w, r.URL.Query().Encode())
}

func init() {
	//注册handler
	apiHandlers = make(map[string]apiHandler)
	apiHandlers["carousel"] = carouselHandler
	apiHandlers["hotauthors"] = hotAuthors
}
