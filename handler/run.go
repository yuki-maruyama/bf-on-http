package handler

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/yuki-maruyama/bf-on-http/util"
	"github.com/yuki-maruyama/brainfxxk/interpreter"
)

var bfTimeout = util.StringToIntWithDefault(util.GetEnv("BF_TIMEOUT", "10"), 10)

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(bfTimeout)*time.Second)
	defer cancel()

	if err := interpreter.Run(ctx, string(input), config); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	io.WriteString(w, string(output.Buffer()))
	defer func() {
		r.Body.Close()
	}()
}
