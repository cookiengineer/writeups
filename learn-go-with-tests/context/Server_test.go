package context

import "context"
import "net/http"
import "net/http/httptest"
import "testing"
import "time"

func TestServer(t *testing.T) {

	data := "Hello, World!"

	t.Run("request context", func(t *testing.T) {

		store := &SpyStore{
			delay:    10 * time.Millisecond,
			response: data,
		}

		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("Expected %s but got %s", data, response.Body.String())
		}

	})

	t.Run("request cancelled context", func(t *testing.T) {

		store := &SpyStore{
			delay:    10 * time.Millisecond,
			response: data,
		}

		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelling_context, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancelling_context)

		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Error("Response should not have been written")
		}

	})

}
