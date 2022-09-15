package services

import (
	"log"
	"github.com/keithyw/pbuf-services/protobufs"
	"golang.org/x/net/context"
)

type Server struct{
	protobufs.UnimplementedUserServer
}

func (s *Server) DeleteUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	log.Printf("ID has been deleted: %d", in.UserId)
	return &protobufs.UserServiceResponse{Msg: "User Deleted"}, nil
}

func (s *Server) SaveUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	log.Printf("Username: %s", *in.Username)
	return &protobufs.UserServiceResponse{Msg: "Successful save!"}, nil
}

func (* Server) UpdateUser(ctx context.Context, in *protobufs.UserMessage) (*protobufs.UserServiceResponse, error) {
	log.Printf("Updated user: %d", in.UserId)
	return &protobufs.UserServiceResponse{Msg: "User updated"}, nil
}