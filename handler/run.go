package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/yuki-maruyama/brainfxxk/interpreter"
)

func RunHandler(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	input := make([]byte, len)
	output := new(bytes.Buffer)
	r.Body.Read(input)

	if err := interpreter.Run(string(input), nil, output); err != nil {
		w.WriteHeader(403)
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, output.String())
	defer func () {
		r.Body.Close()
	}()
}
