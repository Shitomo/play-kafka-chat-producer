package port

import (
	"context"
	"github/Shitomo/producer/domain/model"
)

type UserRegisterInputPort interface {
	Execute(context.Context, UserRegisterInput)
}

type UserRegisterInput struct {
	FirstName model.FirstName
	LastName  model.LastName
	BirthDay  model.BirthDay
}

func NewUserRegisterInput(firstNameStr string, lastNameStr string, birthDayStr string) (UserRegisterInput, error) {
	var errors model.ValidationErrors

	firstName := model.FirstName(firstNameStr)
	if err := firstName.Valid(); err != nil {
		errors = append(errors, err)
	}

	lastName := model.LastName(lastNameStr)
	if err := lastName.Valid(); err != nil {
		errors = append(errors, err)
	}

	birthDay, err := model.NewDatetime(birthDayStr)
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) != 0 {
		return UserRegisterInput{}, errors
	}

	return UserRegisterInput{
		FirstName: firstName,
		LastName:  lastName,
		BirthDay:  model.BirthDay(birthDay),
	}, nil
}

type UserRepository interface {
	Save(ctx context.Context, user model.User) error
}

type UserProducer interface {
	Produce(ctx context.Context, user model.User) error
}

type UserRegisterOutputPort interface {
	Render(context.Context, UserRegisterOutput)
	ErrorRender(context.Context, error)
}

type UserRegisterOutput struct {
	FirstName model.FirstName
	LastName  model.LastName
	BirthDay  model.BirthDay
}
