package concurrency

import "testing"
import "time"

func stubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)

	for u := 0; u < len(urls); u++ {
		urls[u] = "http://localhost:3000/"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(stubWebsiteChecker, urls)
	}

}
