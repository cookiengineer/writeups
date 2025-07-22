package workers

import "context"
import "fmt"
import "sort"
import "sync"

type Scheduler struct {
	Context   context.Context
	Cancel    context.CancelFunc
	Resources map[string]*sync.Mutex
	Statuses  map[string]*Status
	Workers   map[string]*Worker
	channel   chan Status
	mutex     *sync.RWMutex
	waitgroup *sync.WaitGroup
}

func NewScheduler() *Scheduler {

	var scheduler Scheduler

	ctx, cancel := context.WithCancel(context.Background())

	scheduler.Context = ctx
	scheduler.Cancel  = cancel
	scheduler.Resources = make(map[string]*sync.Mutex)
	scheduler.Statuses = make(map[string]*Status)
	scheduler.Workers = make(map[string]*Worker)
	scheduler.channel = make(chan Status)
	scheduler.mutex = &sync.RWMutex{}
	scheduler.waitgroup = &sync.WaitGroup{}

	return &scheduler

}

func (scheduler *Scheduler) CreateWorker(name string, size uint) *Worker {

	worker := Worker{
		Name:    name,
		Context: scheduler.Context,
		Mutex:   scheduler.GetMutexForResource(name),
		Size:    size,
	}

	return &worker

}
func (scheduler *Scheduler) AddWorker(worker *Worker) {

	scheduler.waitgroup.Add(1)
	scheduler.Workers[worker.Name] = worker

}

// TODO: This should contain structs.Cache and structs.Database later
type Callback func()

func (scheduler *Scheduler) AddFunc(callback Callback) {

	// TODO Add a worker instance
	// TODO: Set worker.Callback
	// TODO: Set worker.Context
	// TODO: Set worker.Channel

}

func (scheduler *Scheduler) Start() {

	for _, worker := range scheduler.Workers {

		go worker.Run(
			scheduler.channel,
			scheduler.waitgroup,
		)

	}

	go func() {

		for {
			select {
			case <- scheduler.Context.Done():
				return
			case worker_status := <- scheduler.channel:
				scheduler.mutex.Lock()
				scheduler.Statuses[worker_status.Name] = &worker_status
				scheduler.mutex.Unlock()
			}
		}

	}()

}

func (scheduler *Scheduler) Stop() {
	scheduler.Cancel()
	scheduler.waitgroup.Wait()
	close(scheduler.channel)
}

func (scheduler *Scheduler) GetMutexForResource(name string) *sync.Mutex {

	scheduler.mutex.Lock()

	mutex, ok := scheduler.Resources[name]

	if ok == false {
		scheduler.Resources[name] = &sync.Mutex{}
		mutex = scheduler.Resources[name]
	}

	scheduler.mutex.Unlock()

	return mutex

}

func (scheduler *Scheduler) Print() {

	scheduler.mutex.RLock()

	names := make([]string, 0)

	for name, _ := range scheduler.Statuses {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {

		status := scheduler.Statuses[name]

		fmt.Printf("Worker %s: %s ... %d of %d \n", name, status.Message, status.Current, status.Size)

	}

	scheduler.mutex.RUnlock()

}

func (scheduler *Scheduler) Wait() {
	scheduler.waitgroup.Wait()
	scheduler.Cancel()
}
