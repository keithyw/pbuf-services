package services

import (
	"log"
	"github.com/keithyw/pbuf-services/protobufs"
	"golang.org/x/net/context"
)

type Server struct{
	protobufs.UnimplementedUserServer
}

func (s *Server) SaveUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserMessage, error) {
	log.Printf("Username: %s", in.Username)
	return &protobufs.UserMessage{Username: "Successful save!"}, nil
}