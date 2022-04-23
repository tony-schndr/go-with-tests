package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	CountDown(buffer, sleeper)
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if sleeper.Calls != 4 {
		t.Errorf("got %d want 4", sleeper.Calls)
	}

}

func TestCountDownOrderOfOperations(t *testing.T) {
	spy := &SpyCountdownOperations{}
	CountDown(spy, spy)

	want := []string{
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
		sleep,
		write,
	}
	if !reflect.DeepEqual(spy.Calls, want) {
		t.Errorf("wanted %v got %v", want, spy.Calls)
	}
}

func TestConfigSleeper(t *testing.T) {
	sleepTime := time.Second * 5
	spyTime := &SpyTime{}

	sleeper := ConfigSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
