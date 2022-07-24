package gateway

import (
	"context"
	"github/Shitomo/producer/domain/model"
	"github/Shitomo/producer/driver/db"
	logger "github/Shitomo/producer/driver/log"
	"github/Shitomo/producer/ent"
	"github/Shitomo/producer/usecase/port"
	"time"
)

type UserGateway struct {
	dbClient ent.Client
}

func NewUserGateway() port.UserRepository {
	dbClient, err := db.NewClient()
	if err != nil {
		logger.Log().Sugar().Fatalf("Fail to create db client. caused by", err)
	}
	return UserGateway{
		dbClient: *dbClient,
	}
}

func (u UserGateway) Save(ctx context.Context, user model.User) error {
	birthDayDatetime := model.Datetime(user.BirthDay)
	return u.dbClient.User.Create().
		SetFirstName(string(user.FirstName)).
		SetLastName(string(user.LastName)).
		SetBirthday(time.Time(birthDayDatetime)).
		Exec(ctx)
}
