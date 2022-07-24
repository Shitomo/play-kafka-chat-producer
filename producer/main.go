package main

import (
	"github/Shitomo/producer/domain/model"
	logger "github/Shitomo/producer/driver/log"
	"os"
)

func main() {
	logger.InitLogger()
	logger.Log().Sugar().Infof("%s server starting...", os.Getenv("ENV"))

	model.LoadEnv("config/")
	model.LoadLocation()

	server := InitializeHTTPServer()

	server.Run()
}
