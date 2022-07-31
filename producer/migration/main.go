package main

import (
	"context"
	"github/Shitomo/my-chat/driver/db"
	"github/Shitomo/my-chat/driver/logger"
	"github/Shitomo/my-chat/model"

	_ "github.com/lib/pq"
)

func main() {
	model.LoadEnv("../config/")
	client, err := db.NewClient()
	if err != nil {
		logger.Fatalf(context.Background(), "Failed to create db cient, cause by %s", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logger.Fatalf(ctx, "Failed creating schema resources: %v", err)
	}
}
