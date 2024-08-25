package selects

import "net/http"
import "net/http/httptest"
import "testing"
import "time"

func createTestServer(delay time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}

func TestRacer(t *testing.T) {

	t.Run("Compare response times of two test servers", func(t *testing.T) {

		slow_server := createTestServer(100 * time.Millisecond)
		fast_server := createTestServer(50 * time.Millisecond)

		defer slow_server.Close()
		defer fast_server.Close()

		slow_url := slow_server.URL
		fast_url := fast_server.URL

		want := fast_url
		got, err := Racer(slow_url, fast_url)

		if err != nil {
			t.Fatalf("Unexpected error %v", err)
		}

		if got != want {
			t.Errorf("Expected %q but got %q", want, got)
		}

	})

	t.Run("Compare response times of google and ebay", func(t *testing.T) {

		slow_url := "https://linkedin.com/"
		fast_url := "https://ebay.com/"

		want := fast_url
		got, err := Racer(slow_url, fast_url)

		if err != nil {
			t.Fatalf("Unexpected error %v", err)
		}

		if got != want {
			t.Errorf("Expected %q but got %q", want, got)
		}

	})

	t.Run("Expect error on response timeout", func(t *testing.T) {

		server := createTestServer(200 * time.Millisecond)

		defer server.Close()

		_, err := WebsiteRacer(server.URL, server.URL, 100 * time.Millisecond)

		if err == nil {
			t.Error("Expected timeout error but didn't get one")
		}

	})

}

