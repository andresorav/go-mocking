package countdown

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const countdownStart = 3
const finalWorld = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		printWord(out, strconv.Itoa(i), sleeper)
	}

	printWord(out, finalWorld, sleeper)
}

func printWord(out io.Writer, word string, sleeper Sleeper) {
	sleeper.Sleep()
	fmt.Fprintln(out, word)
}
