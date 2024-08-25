package mocking

import "bytes"
import "reflect"
import "strings"
import "testing"

type SpyCountdown struct {
	Callstack []string
}

func (countdown *SpyCountdown) Sleep() {
	countdown.Callstack = append(countdown.Callstack, "sleep")
}

func (countdown *SpyCountdown) Write(p []byte) (n int, err error) {
	countdown.Callstack = append(countdown.Callstack, "write")
	return 0, nil
}

func TestCountdown(t *testing.T) {

	t.Run("Countdown 10 to Go!", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdown{}, 10)

		got := buffer.String()
		want := strings.Join([]string{
			"10", "9", "8", "7", "6", "5", "4", "3", "2", "1", "Go!",
		}, "\n") + "\n"

		if got != want {
			t.Errorf("Expected %q but got %q", want, got)
		}

	})

	t.Run("Sleep before every Countdown", func(t *testing.T) {

		spycountdown := &SpyCountdown{}
		Countdown(spycountdown, spycountdown, 3)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, spycountdown.Callstack) {
			t.Errorf("Expected callstack %v but got %v", want, spycountdown.Callstack)
		}

	})

}

