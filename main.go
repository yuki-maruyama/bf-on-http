package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/yuki-maruyama/bf-on-http/router"
)

type Config struct {
	Port int
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := router.NewRouter()

	config := &Config{
		Port: 8080,
	}
	server := &http.Server{
		Addr: ":"+fmt.Sprint(config.Port),
		Handler: r,
	}

	go func ()  {
		<- ctx.Done()
		server.Shutdown(ctx)
	}()
	log.Printf("server start running at :%d", config.Port)
	log.Fatal(server.ListenAndServe())
}
