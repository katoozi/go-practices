package main

import (
	"context"
	"fmt"

	"github.com/katoozi/go-practices/gRPC/grpc/messageService"
	"google.golang.org/grpc"
)

var port = ":8080"

// AboutToSayIt will say the text string
func AboutToSayIt(ctx context.Context, m messageService.MessageServiceClient, text string) (*messageService.Response, error) {
	request := &messageService.Request{
		Text:    text,
		Subtext: "New Message",
	}
	r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}

	client := messageService.NewMessageServiceClient(conn)
	r, err := AboutToSayIt(context.Background(), client, "Me Message")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Text:", r.GetText())
	fmt.Println("Response SubText:", r.GetSubtext())
}
