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
		taskUsecase usecases.ITaskUsecase
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
	return c.usecase.taskUsecase.UpdateTaskExpiryTime(ctx)
}

func NewCron(u UsecaseDependency) *Cron {
	return &Cron{
		scheduler: cron.New(),
		usecase:   u,
	}
}

var _ CronService = &Cron{}
