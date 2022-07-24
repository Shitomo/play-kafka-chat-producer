package main

import (
	"context"
	"github/Shitomo/producer/domain/model"
	"github/Shitomo/producer/driver/db"
	logger "github/Shitomo/producer/driver/log"

	_ "github.com/lib/pq"
)

func main() {
	model.LoadEnv("../config/")
	logger.InitLogger()
	client, err := db.NewClient()
	if err != nil {
		logger.Log().Sugar().Fatalf("Failed to create db cient, cause by %s", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logger.Log().Sugar().Fatalf("Failed creating schema resources: %v", err)
	}
}
