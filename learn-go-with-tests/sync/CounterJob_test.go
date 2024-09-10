package sync

import "testing"

func assertCounterJobState(t testing.TB, got CounterJobState, want CounterJobState) {

	t.Helper()

	if got != want {
		t.Errorf("Expected %d to be %d", got, want)
	}

}

func TestCounterJobPool(t *testing.T) {

	t.Run("Create one job", func(t *testing.T) {

		job := CounterJobPool.Get().(*CounterJob)

		job.counter = NewCounter()
		job.limit = 13
		job.Run()

		// Reuse counterjob
		job.state = CounterJobStateRecycled
		CounterJobPool.Put(job)

		assertCounter(t, job.counter, 13)

		job2 := CounterJobPool.Get().(*CounterJob)
		job2.limit = 37

	})

	t.Run("Create and reuse job from pool", func(t *testing.T) {

		job := CounterJobPool.Get().(*CounterJob)

		job.counter = NewCounter()
		job.limit = 13
		job.Run()

		// Reuse counterjob
		job.state = CounterJobStateRecycled
		CounterJobPool.Put(job)

		job2 := CounterJobPool.Get().(*CounterJob)

		assertCounterJobState(t, job2.state, CounterJobStateRecycled)

		job2.limit = 37
		job2.Run()

		assertCounter(t, job.counter, 13+37)

	})

}
