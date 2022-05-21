package countdown

import (
	"fmt"
	"io"
	"time"
)

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type ConfigSleeper struct {
	Duration time.Duration
	Tsleep   func(time.Duration)
}

func (c *ConfigSleeper) Sleep() {
	c.Tsleep(c.Duration)

}

func CountDown(writer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(writer, i)
	}
	s.Sleep()
	fmt.Fprintf(writer, "Go!")
}
