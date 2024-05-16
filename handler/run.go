package handler

import (
	"io"
	"net/http"

	"github.com/yuki-maruyama/bf-on-http/util"
	"github.com/yuki-maruyama/brainfxxk/interpreter"
)

func RunHandler(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	input := make([]byte, len)
	output := util.NewFixedWriter(1024 * 1024)
	r.Body.Read(input)

	config := interpreter.Config {
		MemorySize: 16384,
		MaxStep: 100000000000,

		Reader: nil,
		Writer: output,
	}

	if err := interpreter.Run(string(input), config); err != nil {
		w.WriteHeader(403)
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, string(output.Buffer()))
	defer func () {
		r.Body.Close()
	}()
}
