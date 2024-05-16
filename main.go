package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/yuki-maruyama/bf-on-http/config"
	"github.com/yuki-maruyama/bf-on-http/router"
	"github.com/yuki-maruyama/bf-on-http/util"
)

var portNum = util.StringToIntWithDefault(util.GetEnv("PORT", "8080"), 8080)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	r := router.NewRouter()
	config := &config.Config{
		Port: portNum,
	}
	server := &http.Server{
		Addr: ":"+fmt.Sprint(config.Port),
		Handler: r,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		<- ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()
	log.Printf("server start running at :%d", config.Port)
	log.Fatal(server.ListenAndServe())
	wg.Wait()
}
