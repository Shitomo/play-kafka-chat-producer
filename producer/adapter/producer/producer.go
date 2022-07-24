package producer

import (
	"context"
	"encoding/json"
	"github/Shitomo/producer/domain/model"
	"github/Shitomo/producer/driver/kafka"
	logger "github/Shitomo/producer/driver/log"
	"github/Shitomo/producer/usecase/port"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

type UserProducer struct {
	asyncProducer sarama.AsyncProducer
}

func NewUserProducer() port.UserProducer {
	producer, err := kafka.NewProducer()
	if err != nil {
		logger.Log().Sugar().Fatalf("Fail to create producer. caused by", err)
	}
	return UserProducer{
		asyncProducer: producer,
	}
}

// 送信メッセージ
type UserSendMessage struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDay  int64  `json:"birthDay"`
}

func (u UserProducer) Produce(ctx context.Context, user model.User) error {

	timestamp := time.Now().UnixNano()

	send := &UserSendMessage{
		FirstName: string(user.FirstName),
		LastName:  string(user.LastName),
		BirthDay:  model.Datetime(user.BirthDay).UnixMilli(),
	}

	jsBytes, err := json.Marshal(send)
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "user-topic",
		Key:   sarama.StringEncoder(strconv.FormatInt(timestamp, 10)),
		Value: sarama.StringEncoder(string(jsBytes)),
	}

	u.asyncProducer.Input() <- msg

	select {
	case <-u.asyncProducer.Successes():
		return nil
	case err := <-u.asyncProducer.Errors():
		return err
	case <-ctx.Done():
		return nil
	}
}
