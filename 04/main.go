package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}
}

const EVENT_TIME_FORMAT string = "2006-01-02 15:04"

type EventType uint8

const TYPE_BEGIN EventType = 0
const TYPE_SLEEP EventType = 1
const TYPE_WAKE EventType = 2

type Event interface {
	Time() time.Time
	Guard() int
	Type() EventType
}

func parseLine(line string) Event {
	return nil
}

func sortEvents(evts []Event) []Event {
	return evts
}

type Nap struct {
	Start    time.Time
	Duration int
}

type SleepyGuard struct {
	ID     int
	Asleep []Nap
}

func findSleepyGuards(evts []Event) []SleepyGuard {
	return []SleepyGuard{}
}

func sleepiestGuard(guards []SleepyGuard) SleepyGuard {
	return guards[0]
}

func sleepiestMinute(naps []Nap) int {
	return 0
}
