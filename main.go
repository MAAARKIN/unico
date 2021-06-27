package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/MAAARKIN/unico/api"
	"github.com/MAAARKIN/unico/config"
	"github.com/MAAARKIN/unico/container"
)

func main() {

	cfg := config.Properties()
	cdi := container.Injector(cfg)
	api.StartHttpServer(cfg, cdi)

	ctx := onSignal()
	log.Printf("Server is shutting down..")
	api.Server.Shutdown(ctx)
}

func onSignal() context.Context {
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return ctx
}
