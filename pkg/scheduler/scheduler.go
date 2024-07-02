package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"task_mission/enums"
	"time"
)

type (
	Scheduler struct {
		scheduler gocron.Scheduler
	}
)

var _ IScheduler = &Scheduler{}

func NewScheduler() (*Scheduler, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}
	return &Scheduler{
		scheduler: scheduler,
	}, nil
}

func (s *Scheduler) Task(duration int, timeIn enums.TimeDuration, taskFunction func()) {
	s.scheduler.Start()
	var ms int
	switch timeIn {
	case enums.YEAR:
		ms = 31536000000
	case enums.MONTH:
		ms = 0
	case enums.DAY:
		ms = 86400000
	case enums.HOUR:
		ms = 3600000
	case enums.MINUTE:
		ms = 60000
	case enums.SECOND:
		ms = 1000
	case enums.MILLISECOND:
		ms = 1
	}
	duration = duration * ms
	s.scheduler.NewJob(
		gocron.DurationJob(time.Duration(duration)*time.Millisecond),
		gocron.NewTask(taskFunction),
	)
}

func (s *Scheduler) Start() {
	s.scheduler.Start()
}

func (s *Scheduler) Shutdown() error {
	return s.scheduler.Shutdown()
}
