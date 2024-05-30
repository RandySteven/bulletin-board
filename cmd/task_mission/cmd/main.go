package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"task_mission/apps"
	"task_mission/pkg/config"
	"task_mission/pkg/db"
)

//func init() {
//	err := godotenv.Load("../files/env/.env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//		return
//	}
//}

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

	apps.RegisterMiddleware(r)

	handlers.InitRouter(r)
	config.Run(r)
}
