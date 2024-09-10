package context

import "errors"
import "net/http"

type SpyResponseWriter struct {
	written bool
}

func (writer *SpyResponseWriter) Header() http.Header {
	writer.written = true
	return nil
}

func (writer *SpyResponseWriter) Write([]byte) (int, error) {
	writer.written = true
	return 0, errors.New("Not implemented")
}

func (writer *SpyResponseWriter) WriteHeader(status int) {
	writer.written = true
}
