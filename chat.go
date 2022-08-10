package chatapi

import (
	"context"
	"log"

	chat "github.com/Shitomo/play-kafka-chat-producer/gen/chat"

	"github.com/Shitomo/play-kafka-chat-core/adapter/gateway"
	"github.com/Shitomo/play-kafka-chat-core/driver/logger"
	"github.com/Shitomo/play-kafka-chat-core/model"
)

// chat service example implementation.
// The example methods log the requests and return zero values.
type chatsrvc struct {
	messageGateway          gateway.MessageGateway
	realtimeMessageProducer gateway.RealtimeMessagePublisher
}

// NewChat returns the chat service implementation.
func NewChat(logger *log.Logger, messageGateway gateway.MessageGateway, realtimeMessageProducer gateway.RealtimeMessagePublisher) chat.Service {
	return &chatsrvc{
		messageGateway:          messageGateway,
		realtimeMessageProducer: realtimeMessageProducer,
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
	err = s.realtimeMessageProducer.Produce(ctx, message)
	if err != nil {
		logger.Errorf(ctx, "Failed to produce message to realtime message server. caused by %v", err)
		return nil, err
	}

	return res, nil
}
