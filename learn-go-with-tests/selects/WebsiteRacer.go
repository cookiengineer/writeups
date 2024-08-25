package selects

import "fmt"
import "net/http"
import "time"

func doRequest(url string) chan struct{} {

	channel := make(chan struct{})

	go func() {
		http.Get(url)
		close(channel)
	}()

	return channel

}

func WebsiteRacer(url_a string, url_b string, timeout time.Duration) (string, error) {

	select {
	case <-doRequest(url_a):
		return url_a, nil
	case <-doRequest(url_b):
		return url_b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timeout after waiting for both %s and %s", url_a, url_b)
	}

}

