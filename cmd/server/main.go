package main

import (
	"log"
	"net"

	usersService "github.com/pstano1/go-cart/internal/users"
	"github.com/pstano1/go-cart/protopb/users"
	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	usersService := usersService.NewService()
	users.RegisterUsersServiceServer(server, &usersService)

	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
