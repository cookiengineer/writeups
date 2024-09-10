package sync

import "sync"
import "testing"

func assertCounter(t testing.TB, counter *Counter, want int) {

	t.Helper()

	got := counter.Value()

	if got != want {
		t.Errorf("Expected %d to be %d", got, want)
	}

}

func TestCounter(t *testing.T) {

	t.Run("Increment to 3 synchronously", func(t *testing.T) {

		counter := NewCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)

	})

	t.Run("Increment to 1000 concurrently", func(t *testing.T) {

		counter := NewCounter()

		var wg sync.WaitGroup

		wg.Add(1000)

		for i := 0; i < 1000; i++ {

			go func() {
				counter.Increment()
				wg.Done()
			}()

		}

		wg.Wait()

		assertCounter(t, counter, 1000)

	})

}

