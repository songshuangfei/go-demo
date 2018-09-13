package handler

import (
	"io"
	"net/http"
	"os"
)

//UploadHandler ...
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //get请求发送页面
		if err := renderHTML(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if r.Method == "POST" { //post请求接受数据
		f, h, err := r.FormFile("image") //获取上传来的表单的图片
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //返回服务器内部错误级原因
			return
		}

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(uploadDir + "/" + filename) //新建一个图片文件
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil { //图片复制到新建的文件
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, //重定向到刚上传的文件
			http.StatusFound)
	}
}
