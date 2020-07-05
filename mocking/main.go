package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper is an implementation of Sleeper with a predefined delay.
type DefaultSleeper struct {
}

// Sleep will pause execution for the defined Duration.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

const finalWord = "  >> Go!"
const countdownstart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownstart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
