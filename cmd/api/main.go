package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/matherique/lp2-sbo-api/pkg/config"
	server "github.com/matherique/lp2-sbo-api/pkg/http"
	"github.com/matherique/lp2-sbo-api/pkg/store"
	"github.com/matherique/lp2-sbo-api/pkg/utils"
)

func main() {
	config, err := config.Read()

	log := utils.NewLogger("APP")

	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	store := store.NewStore(
		config.CONNECTION_STRING,
	)

	if err := store.Connect(); err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(config.APP_PORT)

	loadRoutes(srv, store)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		store.Close()
		cancel()
	}()

	if err := srv.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
