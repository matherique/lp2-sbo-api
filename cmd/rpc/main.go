package main

import (
	"log"
	"net"

	"github.com/matherique/lp2-sbo-api/internal/app"
	proto "github.com/matherique/lp2-sbo-api/internal/proto/book"
	"github.com/matherique/lp2-sbo-api/pkg/config"
	"github.com/matherique/lp2-sbo-api/pkg/store"
	"github.com/matherique/lp2-sbo-api/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	log := utils.NewLogger("RPC")
	lis, err := net.Listen("tcp", config.RPC_PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	store := store.NewStore(
		config.CONNECTION_STRING,
	)

	s := grpc.NewServer()
	srv := app.NewBookService(store)

	proto.RegisterBookServiceServer(s, srv)

	log.Println("start server")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
