package main

import (
	"context"
	"fmt"
	"net"

	"github.com/katoozi/go-practices/gRPC/grpc/messageService"
	"google.golang.org/grpc"
)

var port = ":8080"

// MessageService will implement MessageService interface
type MessageService struct {
}

// SayIt will trigger the SayIt rpc call
func (MessageService) SayIt(ctx context.Context, r *messageService.Request) (*messageService.Response, error) {
	fmt.Println("Request Text:", r.GetText())
	fmt.Println("Request SubText:", r.GetSubtext())

	response := &messageService.Response{
		Text:    r.GetText(),
		Subtext: "Got it!",
	}

	return response, nil
}

func main() {
	server := grpc.NewServer()
	var messageServer MessageService

	messageService.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")
	server.Serve(listen)
}
