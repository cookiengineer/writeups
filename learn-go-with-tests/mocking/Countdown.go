package mocking

import "fmt"
import "io"

func Countdown(output io.Writer, sleeper Sleeper, iterations int) {

	for i := iterations; i > 0; i-- {
		fmt.Fprintln(output, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(output, "Go!")

}
