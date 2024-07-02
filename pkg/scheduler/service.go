package scheduler

import (
	"context"
	"task_mission/enums"
)

type (
	IScheduler interface {
		Task(duration int, timeIn enums.TimeDuration, taskFunction func())
		Start()
		Shutdown() error
	}

	CronService interface {
		RunAllJob(ctx context.Context) error
		autoUpdateExpiryTime(ctx context.Context) error
	}
)
