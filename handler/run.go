package handler

import (
	"io"
	"net/http"

	"github.com/yuki-maruyama/bf-on-http/util"
	"github.com/yuki-maruyama/brainfxxk/interpreter"
)

func RunHandler(w http.ResponseWriter, r *http.Request) {
	var input []byte
	output := util.NewFixedWriter(1024 * 1024)
	for {
		buf := make([]byte, 1024)
		n, err := r.Body.Read(buf)
		if err != nil && err != io.EOF {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		input = append(input, buf[:n]...)
		if err == io.EOF {
			break
		}
	}

	config := &interpreter.Config{
		MemorySize: 16384,
		MaxStep:    100000000000,

		Reader: nil,
		Writer: output,
	}

	if err := interpreter.Run(r.Context(), string(input), config); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	io.WriteString(w, string(output.Buffer()))
	defer func() {
		r.Body.Close()
	}()
}
