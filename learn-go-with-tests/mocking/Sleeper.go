package mocking

import "time"

type Sleeper interface {
	Sleep()
}

type TimeSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (sleeper *TimeSleeper) Sleep() {
	sleeper.sleep(sleeper.duration)
}

