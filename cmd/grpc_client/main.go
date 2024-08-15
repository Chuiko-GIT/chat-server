package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Chuiko-GIT/chat-server/pkg/chat_api"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	defer func() {
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
	}()

	c := chat_api.NewChatApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createChat, err := c.Create(ctx, &chat_api.CreateRequest{})
	if err != nil {
		log.Fatalf("failed to create chat: %v", err)
	}

	log.Printf(color.RedString("Chat create:\n"), color.GreenString("%+v", createChat.GetId()))

	sendMessageChat, err := c.SendMessage(ctx, &chat_api.SendMessageRequest{})
	if err != nil {
		log.Fatalf("failed to send message chat: %v", err)
	}

	log.Printf(color.RedString("Chat send message:\n"), color.GreenString("%+v", sendMessageChat.String()))

	deleteChat, err := c.Delete(ctx, &chat_api.DeleteRequest{})
	if err != nil {
		log.Fatalf("failed to delete chat: %v", err)
	}

	log.Printf(color.RedString("Chat delete:\n"), color.GreenString("%+v", deleteChat.String()))
}
