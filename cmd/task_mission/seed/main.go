package main

import (
	"context"
	"log"
	"task_mission/pkg/config"
	"task_mission/pkg/db"
)

func main() {
	configPath, err := config.ParseFlags()
	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories, err := db.NewRepositories(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories.SeedRole(ctx)
	repositories.SeedCategory(ctx)
}
