package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/keithyw/pbuf-services/protobufs"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting %s", err)
	}
	defer conn.Close()

	client := protobufs.NewUserClient(conn)
	res, err := client.SaveUser(context.Background(), &protobufs.UserMessage{Username: "testuser"})
	if err != nil {
		log.Fatalf("Failed calling UserMessage: %s", err)
	}
	log.Printf("Response: %s", res.Username)
}