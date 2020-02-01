package countdown

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
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
