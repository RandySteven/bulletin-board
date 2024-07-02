package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"task_mission/apps"
	"task_mission/pkg/config"
	"task_mission/pkg/db"
	scheduler2 "task_mission/pkg/scheduler"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	services, err := apps.NewServices(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	//shutdown, err := grafana.SetupOTelSDK(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//defer shutdown(ctx)

	cron := scheduler2.NewCron(apps.NewSchedulers(repositories))
	err = cron.RunAllJob(ctx)
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
