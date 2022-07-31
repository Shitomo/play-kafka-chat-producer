package chatapi

import (
	chat "/Users/shikimeiasaakira/ws/go/fs-apache-kafka/gen/chat"
	"context"
	"log"
)

// chat service example implementation.
// The example methods log the requests and return zero values.
type chatsrvc struct {
	logger *log.Logger
}

// NewChat returns the chat service implementation.
func NewChat(logger *log.Logger) chat.Service {
	return &chatsrvc{logger}
}

// SendMessage implements SendMessage.
func (s *chatsrvc) SendMessage(ctx context.Context, p *chat.SendMessageRequestBody) (res *chat.SendMessageResponseBody, err error) {
	res = &chat.SendMessageResponseBody{}
	s.logger.Print("chat.SendMessage")
	return
}
