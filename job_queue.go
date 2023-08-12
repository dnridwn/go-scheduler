package goscheduler

import (
	"context"
	"runtime"
)

type JobQueue struct {
	jobs chan Schedule
}

func NewJobQueue() JobQueue {
	return JobQueue{
		jobs: make(chan Schedule, runtime.GOMAXPROCS(-1)),
	}
}

func (q *JobQueue) Add(job *Schedule) {
	q.jobs <- *job
}

func (q *JobQueue) Listen(ctx context.Context) {
	go func() {
		for job := range q.jobs {
			job.job()
		}
	}()
}

func (q *JobQueue) Stop() {
	close(q.jobs)
}
