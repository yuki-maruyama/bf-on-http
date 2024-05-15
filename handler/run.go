package handler

import (
	"io"
	"net/http"
)

func RunHandler(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	io.WriteString(w, string(body))
	defer func () {
		r.Body.Close()
	}()
}
