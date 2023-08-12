package goscheduler

import (
	"errors"
	"time"

	"github.com/adhocore/gronx"
	"github.com/google/uuid"
)

func NewJobId() string {
	return uuid.New().String()
}

type Schedule struct {
	JobId    string
	cronExpr string
	job      func()
}

type Scheduler struct {
	schedules []Schedule
	gronx     gronx.Gronx
	jobQueue  JobQueue
}

func NewScheduler() Scheduler {
	return Scheduler{
		schedules: make([]Schedule, 0),
		gronx:     gronx.New(),
		jobQueue:  NewJobQueue(),
	}
}

func (s *Scheduler) Add(cronExpr string, job func()) error {
	if !s.gronx.IsValid(cronExpr) {
		return errors.New("invalid cron expression")
	}

	s.schedules = append(s.schedules, Schedule{
		JobId:    NewJobId(),
		cronExpr: cronExpr,
		job:      job,
	})

	return nil
}

func (s *Scheduler) Run() {
	for _, sch := range s.schedules {
		go func(sch Schedule) {
			ticker := time.NewTicker(time.Second)
			for {
				select {
				case <-ticker.C:
					if ok, err := s.gronx.IsDue(sch.cronExpr); ok && err == nil {
						s.jobQueue.Add(&sch)
					}
				}
			}
		}(sch)
	}

	s.jobQueue.Listen()
}
