package goscheduler

import (
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

func (q *JobQueue) Listen() {
	go func() {
		for job := range q.jobs {
			job.job()
		}
	}()
}

func (q *JobQueue) Clear() {
	close(q.jobs)
}
