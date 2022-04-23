package main

import (
	"go-with-tests/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigSleeper{Duration: 5 * time.Second, Tsleep: time.Sleep}
	countdown.CountDown(os.Stdout, sleeper)
}
