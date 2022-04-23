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

type SpyCountdownOperations struct {
	Calls []string
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {
}

type ConfigSleeper struct {
	Duration time.Duration
	Tsleep   func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigSleeper) Sleep() {
	c.Tsleep(c.Duration)

}

func (s *SpyTime) Sleep(duration time.Duration) {

	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls += 1
}

func CountDown(writer io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(writer, i)
	}
	s.Sleep()
	fmt.Fprintf(writer, "Go!")
}
