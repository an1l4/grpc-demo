package main

import (
	"log"
	"net"

	"github.com/an1l4/grpc-demo/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000:%v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer,&s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000:%v", err)
	}
}
