package handler

import (
	"net/http"
)

//ViewHandler ...
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := uploadDir + "/" + imageID      //拼接查询请求的图片路径
	if exists := isExists(imagePath); !exists { //判断图片是否存在
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-type", "image") //返回找到的图片文件
	http.ServeFile(w, r, imagePath)
}
