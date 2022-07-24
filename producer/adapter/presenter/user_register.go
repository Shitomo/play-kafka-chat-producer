package presenter

import (
	"context"
	"github/Shitomo/producer/adapter"
	"github/Shitomo/producer/usecase/port"
	"net/http"
)

type UserRegisterPresenter struct{}

func NewUserRegisterPresenter() port.UserRegisterOutputPort {
	return &UserRegisterPresenter{}
}

func (p UserRegisterPresenter) Render(ctx context.Context, output port.UserRegisterOutput) {
	r, _ := adapter.GetResWriter(ctx)

	r.WriteHeader(http.StatusOK)
}

func (p UserRegisterPresenter) ErrorRender(ctx context.Context, err error) {
	ErrorRender(ctx, err)
}
