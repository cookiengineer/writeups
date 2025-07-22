package workers

import "context"
import "sync"
import "time"

type Worker struct {
	Name    string
	Context context.Context
	Mutex   *sync.Mutex
	Size    uint
}

func NewWorker(name string, ctx context.Context, mutex *sync.Mutex) *Worker {

	var worker Worker

	worker.Name = name
	worker.Context = ctx
	worker.Mutex = mutex

	return &worker

}

func (worker Worker) Run(channel chan<-Status, waitgroup *sync.WaitGroup) {

	worker.Mutex.Lock()

	for i := uint(0); i < worker.Size; i++ {

		select {
		case <-worker.Context.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)

			channel <- Status{
				Name: worker.Name,
				Size: worker.Size,
				Current: i+1,
				Message: "Crunching numbers",
			}

		}

	}

	worker.Mutex.Unlock()
	waitgroup.Done()

}
