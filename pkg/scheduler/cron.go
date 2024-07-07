package scheduler

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task_mission/interfaces/usecases"
	"time"
)

type (
	Cron struct {
		scheduler *cron.Cron
		usecase   UsecaseDependency
	}

	UsecaseDependency struct {
		TaskUsecase usecases.ITaskUsecase
	}
)

func (c *Cron) RunAllJob(ctx context.Context) (err error) {
	defer c.scheduler.Stop()
	err = c.autoUpdateExpiryTime(ctx)
	if err != nil {
		return err
	}
	go c.scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	return nil
}

func (c *Cron) autoUpdateExpiryTime(ctx context.Context) (err error) {
	_, err = c.scheduler.AddFunc("@hourly", func() {
		err = c.usecase.TaskUsecase.UpdateTaskExpiryTime(ctx)
		if err != nil {
			return
		}
		log.Println("this expiry time is ", time.Now())
	})
	return
}

func NewCron(u UsecaseDependency) *Cron {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &Cron{
		scheduler: cron.New(cron.WithLocation(jakartaTime)),
		usecase:   u,
	}
}

var _ CronService = &Cron{}
