//go:generate wire
package main

import (
	"context"
	"github/Shitomo/producer/domain/model"
	"github/Shitomo/producer/driver/logger"
	"os"
)

func main() {
	logger.Infof(context.Background(), "%s server starting...", os.Getenv("ENV"))

	model.LoadEnv("config/")
	model.LoadLocation()

	server := InitializeHTTPServer()

	server.Run()
}
