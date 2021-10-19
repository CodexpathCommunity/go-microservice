package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/api"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/clients"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/config"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service"
)

var (
	configPath *string
	env        *string
)

func init() {
	// Config file path flag
	configPath = flag.String("config", "", "Path to config file")
	// Env flag
	env = flag.String("env", "development", "Environment")

	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.New(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	clients, err := clients.NewClients(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	service := service.New(clients)
	srv := &http.Server{
		Addr:    conf.App.Addr,
		Handler: api.NewServer(service),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Listening on %s", conf.App.Addr)
	<-stop

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
