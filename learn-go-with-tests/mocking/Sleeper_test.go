package mocking

import "testing"
import "time"

type SpyTime struct {
	SleepDuration time.Duration
}

func (self *SpyTime) Sleep(duration time.Duration) {
	self.SleepDuration = duration
}

func TestTimeSleeper(t *testing.T) {

	spytime := &SpyTime{}
	sleeper := TimeSleeper{
		duration: 2 * time.Second,
		sleep:    spytime.Sleep,
	}
	sleeper.Sleep()

	if spytime.SleepDuration != 2 * time.Second {
		t.Errorf("Expected sleep for %v seconds but slept %v", 2 * time.Second, spytime.SleepDuration)
	}

}

