package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	mux  *http.ServeMux
	port string
}

func NewServer(port string, mux *http.ServeMux) *Server {
	s := new(Server)
	s.mux = mux
	s.port = port

	return s
}

func (srv *Server) Start(ctx context.Context, mux *http.ServeMux) error {
	s := http.Server{
		Addr:    srv.port,
		Handler: srv.mux,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("server start in: localhost%s\n", srv.port)

	<-ctx.Done()

	log.Println("server stoped")

	ctxSD, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctxSD); err != nil {
		return err
	}

	log.Printf("server exited properly")

	return nil
}
