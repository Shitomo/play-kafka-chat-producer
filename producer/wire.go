//go:build wireinject
// +build wireinject

package main

import (
	"github/Shitomo/producer/adapter/controller"
	"github/Shitomo/producer/adapter/gateway"
	"github/Shitomo/producer/adapter/messaging"
	"github/Shitomo/producer/adapter/presenter"
	"github/Shitomo/producer/driver/server"
	"github/Shitomo/producer/usecase/interactor"

	"github.com/google/wire"
)

var messagingSet = wire.NewSet(
	messaging.NewUserProducer,
)

var gatewaySet = wire.NewSet(
	gateway.NewUserGateway,
)

var presenterSet = wire.NewSet(
	presenter.NewUserRegisterPresenter,
)

var interactorSet = wire.NewSet(
	interactor.NewUserRegisterInteractor,
)

var controllerSet = wire.NewSet(
	controller.NewUserRegisterController,
)

func InitializeHTTPServer() *server.HTTPServer {
	wire.Build(
		messagingSet,
		gatewaySet,
		presenterSet,
		interactorSet,
		controllerSet,
		server.NewHTTPServer,
	)

	return nil
}
