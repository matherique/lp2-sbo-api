package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	server "github.com/matherique/lp2-sbo-api/pkg/http_server"
	"github.com/matherique/lp2-sbo-api/pkg/store"
)

const port = ":5000"

func main() {
	mux := http.NewServeMux()

	store := store.NewStore()
	if err := store.Connect(); err != nil {
		log.Fatal(err)
	}

	loadRoutes(mux, store)

	srv := server.NewServer(port, mux)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		store.Close()
		cancel()
	}()

	if err := srv.Start(ctx, mux); err != nil {
		log.Fatal(err)
	}
}
