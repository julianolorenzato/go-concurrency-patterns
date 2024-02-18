package patterns

import (
	"fmt"
	"time"
)

type Worker struct {
	ID int
}

func (w Worker) Work(job Job) {
	job.Execute()
}

type Job struct {
	ID int
	time.Duration
}

func (j Job) Execute() {
	time.Sleep(j.Duration)
	fmt.Printf("Job %d finalized\n", j.ID)
}

func WorkerPool() {
	jobs := []Job{{
		ID:       0,
		Duration: time.Second * 2,
	}, {
		ID:       1,
		Duration: time.Second * 4,
	}, {
		ID:       2,
		Duration: time.Second * 5,
	}}

	for i := 0; i < 3; i++ {

	}
}
