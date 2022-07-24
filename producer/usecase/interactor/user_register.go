package interactor

import (
	"context"
	"github/Shitomo/producer/domain/model"
	"github/Shitomo/producer/usecase/port"
)

type UserRegisterInteractor struct {
	userRepository port.UserRepository
	userProducer   port.UserProducer
	output         port.UserRegisterOutputPort
}

func NewUserRegisterInteractor(
	userRepository port.UserRepository,
	userProducer port.UserProducer,
	output port.UserRegisterOutputPort,
) port.UserRegisterInputPort {
	return &UserRegisterInteractor{
		userRepository: userRepository,
		userProducer:   userProducer,
		output:         output,
	}
}

func (i *UserRegisterInteractor) Execute(ctx context.Context, input port.UserRegisterInput) {
	user := model.NewUser(input.FirstName, input.LastName, input.BirthDay)
	err := i.userRepository.Save(ctx, user)
	if err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	if err := i.userProducer.Produce(ctx, user); err != nil {
		i.output.ErrorRender(ctx, err)

		return
	}

	i.output.Render(ctx, port.UserRegisterOutput{})
}
