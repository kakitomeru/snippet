package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/kakitomeru/shared/database"
	"github.com/kakitomeru/shared/env"
	"github.com/kakitomeru/shared/logger"
	. "github.com/kakitomeru/snippet/internal/app"
	"github.com/kakitomeru/snippet/internal/config"
	"github.com/kakitomeru/snippet/internal/model"
)

func main() {
	logger.InitSlog("snippet", "dev", slog.LevelDebug)
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(ctx, "failed to load snippet config", err)
		os.Exit(1)
	}

	if err := env.LoadEnv(cfg.Env); err != nil {
		logger.Error(ctx, "failed to load env", err)
		os.Exit(1)
	}

	logger.Debug(ctx, "Connecting to database")
	db, err := database.ConnectDatabase()
	if err != nil {
		logger.Error(ctx, "failed to connect to database", err)
		os.Exit(1)
	}

	logger.Debug(ctx, "Running migration for table snippet")
	if err = database.Migrate(db, model.Snippet{}); err != nil {
		logger.Error(ctx, "failed to run migration", err)
		os.Exit(1)
	}

	logger.Debug(ctx, "Starting app")
	app := NewApp(db, cfg)

	app.Start(ctx)
}
