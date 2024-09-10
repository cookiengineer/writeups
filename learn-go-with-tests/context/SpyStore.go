package context

import "context"
import "fmt"
import "time"

type SpyStore struct {
	delay    time.Duration
	response string
}

func (store *SpyStore) Fetch(ctx context.Context) (string, error) {

	channel := make(chan string, 1)

	go func() {

		result := ""

		for _, chr := range store.response {

			select {
			case <-ctx.Done():
				fmt.Println("SpyStore cancelled")
				return
			default:
				time.Sleep(store.delay)
				result += string(chr)
			}

		}

		channel <- result

	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-channel:
		return result, nil
	}

}
