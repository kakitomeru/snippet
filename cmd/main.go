package main

import (
	"context"
	"log"

	"github.com/kakitomeru/shared/database"
	"github.com/kakitomeru/shared/env"
	. "github.com/kakitomeru/snippet/internal/app"
	"github.com/kakitomeru/snippet/internal/config"
	"github.com/kakitomeru/snippet/internal/model"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load snippet config: %v", err)
	}

	if err := env.LoadEnv(cfg.Env); err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Running migration for table snippet")
	if err = database.Migrate(db, model.Snippet{}); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Starting app")
	app := NewApp(db, cfg)
	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
