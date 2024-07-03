package scheduler

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
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
	err = c.autoUpdateExpiryTime(ctx)
	if err != nil {
		return err
	}
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
	return &Cron{
		scheduler: cron.New(),
		usecase:   u,
	}
}

var _ CronService = &Cron{}
