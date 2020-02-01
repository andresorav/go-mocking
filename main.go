package main

import "os"
import "./countdown"

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
