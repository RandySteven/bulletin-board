package scheduler

import (
	"context"
	"github.com/robfig/cron/v3"
	"task_mission/interfaces/usecases"
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
	_, err = c.scheduler.AddFunc("@daily", func() {
		err = c.usecase.TaskUsecase.UpdateTaskExpiryTime(ctx)
		if err != nil {
			return
		}
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
