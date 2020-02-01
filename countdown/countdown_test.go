package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *CountdownOperationSpy) Write(_ []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	var spySleeper = &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Not enought calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}

func TestSleepBeforeEveryWrite(t *testing.T) {
	spySleepPrinter := &CountdownOperationSpy{}
	Countdown(spySleepPrinter, spySleepPrinter)

	want := []string{
		"sleep", "write",
		"sleep", "write",
		"sleep", "write",
		"sleep", "write",
	}

	if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
		t.Errorf("Wanted calls %v got %v", want, spySleepPrinter.Calls)
	}
}
