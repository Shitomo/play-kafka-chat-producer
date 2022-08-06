package chatapi

import (
	"context"
	"github/Shitomo/my-chat/adapter/gateway"
	"github/Shitomo/my-chat/driver/logger"
	chat "github/Shitomo/my-chat/gen/chat"
	"github/Shitomo/my-chat/model"
	"log"
)

// chat service example implementation.
// The example methods log the requests and return zero values.
type chatsrvc struct {
	messageGateway gateway.MessageGateway
}

// NewChat returns the chat service implementation.
func NewChat(logger *log.Logger, messageGateway gateway.MessageGateway) chat.Service {
	return &chatsrvc{
		messageGateway: messageGateway,
	}
}

// SendMessage implements SendMessage.
func (s *chatsrvc) SendMessage(ctx context.Context, p *chat.SendMessageRequestBody) (*chat.SendMessageResponseBody, error) {
	res := &chat.SendMessageResponseBody{}
	message := model.NewMessage(*p.SenderID, *p.Content)
	err := s.messageGateway.Save(ctx, message)
	if err != nil {
		logger.Errorf(ctx, "Failed to save message. caused by %v", err)
		return nil, err
	}

	return res, nil
}
