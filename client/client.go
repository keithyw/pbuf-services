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
	username := "testuser"
	res, err := client.SaveUser(context.Background(), &protobufs.UserMessage{UserId: 123, Username: &username})
	if err != nil {
		log.Fatalf("Failed calling SaveUser: %s", err)
	}
	log.Printf("Response: %s", res.Msg)

	res, err = client.UpdateUser(context.Background(), &protobufs.UserMessage{UserId: 123, Username: &username})
	if err != nil {
		log.Fatalf("Failed calling UpdateUser: %s", err)
	}
	log.Printf("Response: %s", res.Msg)

	res, err = client.DeleteUser(context.Background(), &protobufs.UserMessage{UserId: 123})
	if err != nil {
		log.Fatalf("Failed calling Deleting: %s", err)
	}
	log.Printf("Response: %s", res.Msg)
}