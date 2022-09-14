package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/keithyw/pbuf-services/protobufs"
	"github.com/keithyw/pbuf-services/services"
)

func main() {
	fmt.Println("server starting")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed listening: %v", err)
	}

	s := services.Server{}

	serv := grpc.NewServer()
	reflection.Register(serv)
	protobufs.RegisterUserServer(serv, &s)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed serving: %s", err)
	}
}