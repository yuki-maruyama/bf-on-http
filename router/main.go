package router

import (
	"net/http"

	"github.com/yuki-maruyama/bf-on-http/handler"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	mux.HandleFunc("POST /run", handler.RunHandler)

	return mux
}
