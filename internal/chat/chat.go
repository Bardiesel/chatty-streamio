package chat

import (
	"context"
	"log"
	"os"

	stream "github.com/GetStream/stream-chat-go/v6"
)

type Service struct {
	Channel *stream.Channel
}

func NewService() (*Service, error) {
	APIKey := os.Getenv("GETSTREAM_API_KEY")
	APISecret := os.Getenv("GETSTREAM_API_SECRET")

	client, err := stream.NewClient(APIKey, APISecret)
	if err != nil {
		log.Println("Error creating client", err.Error())
		return nil, err
	}

	resp, err := client.CreateChannel(
		context.Background(),
		"messaging",
		"chatty",
		"channel-id",
		&stream.ChannelRequest{
			Members: []string{},
		},
	)
	if err != nil {
		log.Println("Error creating channel", err.Error())
		return nil, err
	}
	log.Printf("Channel created: %+v\n", resp)

	return &Service{
		Channel: resp.Channel,
	}, nil
}

func (s *Service) AddUser(ctx context.Context, userID string) error {
	resp, err := s.Channel.AddMembers(ctx, []string{userID})
	if err != nil {
		log.Println("Error adding user to channel", err.Error())
		return err
	}
	log.Printf("User added to channel: %+v\n", resp)
	return nil
}

func (s *Service) SendMessage(ctx context.Context, userID string, message string) error {
	resp, err := s.Channel.SendMessage(ctx, &stream.Message{
		Text: message,
	}, userID)
	if err != nil {
		log.Println("Error sending message", err.Error())
		return err
	}
	log.Printf("Message sent: %+v\n", resp)
	return nil
}

func (s *Service) PrintMessages() {
	for _, message := range s.Channel.Messages {
		log.Printf("Message: %+v\n", message)
	}
}
