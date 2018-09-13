package handler

import (
	"io/ioutil"
	"net/http"
)

//ListHandler ....
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderData := make(map[string]interface{})
	images := []string{}
	for _, fileinfo := range fileInfoArr {
		imgID := fileinfo.Name()
		images = append(images, imgID)
	}
	renderData["images"] = images
	if err := renderHTML(w, "list", renderData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
