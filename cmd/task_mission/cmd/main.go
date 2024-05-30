package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"task_mission/apps"
	"task_mission/pkg/config"
	"task_mission/pkg/db"
)

func main() {
	configPath, err := config.ParseFlags()

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

	services, err := apps.NewServices(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers := apps.NewHandlers(repositories, services)

	r := mux.NewRouter()

	log.Println("UDAH KE RUN WA")

	r = apps.RegisterMiddleware(r)

	handlers.InitRouter(r)
	config.Run(r)
}
