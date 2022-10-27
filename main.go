package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Bardiesel/chatty-streamio.git/internal/chat"
)

func main() {
	fmt.Println("Hello, playground")

	chatService, err := chat.NewService()
	if err != nil {
		log.Println("Error creating chat service", err.Error())
		return
	}

	ctx := context.Background()

	err = chatService.AddUser(ctx, "user1")
	if err != nil {
		log.Println("Error adding user to channel", err.Error())
		return
	}

	err = chatService.SendMessage(ctx, "user1", "Hello, world!")
	if err != nil {
		log.Println("Error sending message", err.Error())
		return
	}

	chatService.PrintMessages()

}
