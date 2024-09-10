package sync

import "fmt"
import "sync"

type CounterJobState int

const (
	CounterJobStateNew CounterJobState = iota
	CounterJobStateRunning
	CounterJobStateDone
	CounterJobStateRecycled
)

type CounterJob struct {
	counter *Counter
	limit   int
	state   CounterJobState
}

func (job *CounterJob) Run() {

	switch job.state {
	case CounterJobStateNew:
		fmt.Println("Job was allocated")
	case CounterJobStateRecycled:
		fmt.Println("Job was recycled")
	}

	for c := 0; c < job.limit; c++ {
		job.counter.Increment()
	}

	job.state = CounterJobStateDone

}

var CounterJobPool = sync.Pool{
	New: func() interface{} {
		return new(CounterJob)
	},
}

