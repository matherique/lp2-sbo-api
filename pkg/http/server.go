package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/matherique/lp2-sbo-api/pkg/utils"
)

type Server struct {
	mux  *http.ServeMux
	port string
	log  *log.Logger
}

func NewServer(port string, mux *http.ServeMux) *Server {
	s := new(Server)
	s.mux = mux
	s.port = port
	s.log = utils.NewLogger("Server")

	return s
}

func (srv *Server) Start(ctx context.Context, mux *http.ServeMux) error {
	s := http.Server{
		Addr:    srv.port,
		Handler: srv.mux,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			srv.log.Fatal(err)
		}
	}()

	srv.log.Printf("server start in: localhost%s\n", srv.port)

	<-ctx.Done()

	srv.log.Println("server stoped")

	ctxSD, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctxSD); err != nil {
		return err
	}

	srv.log.Printf("server exited properly")

	return nil
}
