package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var events []*Event

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		events = append(events, parseLine(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	events = sortEvents(events)
	shifts := findShifts(events)

	guard_id := sleepiestGuard(shifts)

	naps := []Nap{}
	for _, shift := range shifts {
		if shift.Guard == guard_id {
			naps = append(naps, shift.Naps...)
		}
	}

	minute := sleepiestMinute(naps)

	log.Println("A:", guard_id*minute)

	guard_id, minute, _ = sleepiestMinuteAllTime(sleepsPerMinute(shifts))

	log.Println("B:", guard_id*minute)
}

const EVENT_TIME_FORMAT string = "2006-01-02 15:04"

type EventType uint8

const TYPE_BEGIN EventType = 0
const TYPE_SLEEP EventType = 1
const TYPE_WAKE EventType = 2

var EventMatcher *regexp.Regexp
var BeginMatcher *regexp.Regexp

func init() {
	EventMatcher = regexp.MustCompile("\\[(\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2})\\] (.*)")
	BeginMatcher = regexp.MustCompile("Guard #(\\d+) begins shift")
}

type Event struct {
	Time  time.Time
	Guard int
	Type  EventType
}

func (e *Event) String() string {
	switch e.Type {
	case TYPE_BEGIN:
		return fmt.Sprintf("[%s] Guard #%d begins shift", e.Time.Format(EVENT_TIME_FORMAT), e.Guard)
	case TYPE_SLEEP:
		return fmt.Sprintf("[%s] falls asleep", e.Time.Format(EVENT_TIME_FORMAT))
	case TYPE_WAKE:
		return fmt.Sprintf("[%s] wakes up", e.Time.Format(EVENT_TIME_FORMAT))
	}
	return "UNKNOWN EVENT TYPE"
}

func parseLine(line string) *Event {
	e := Event{Guard: -1}

	m := EventMatcher.FindStringSubmatch(line)
	t, err := time.Parse(EVENT_TIME_FORMAT, m[1])
	if err != nil {
		panic(err)
	}
	e.Time = t

	switch m[2] {
	case "falls asleep":
		e.Type = TYPE_SLEEP
	case "wakes up":
		e.Type = TYPE_WAKE
	default:
		e.Type = TYPE_BEGIN
	}

	if e.Type == TYPE_BEGIN {
		bm := BeginMatcher.FindStringSubmatch(m[2])
		gid, err := strconv.Atoi(bm[1])
		if err != nil {
			panic(err)
		}
		e.Guard = gid
	}

	return &e
}

func sortEvents(evts []*Event) []*Event {
	sort.Slice(evts, func(i, j int) bool { return evts[i].Time.Before(evts[j].Time) })
	return evts
}

type Nap struct {
	Start time.Time
	End   time.Time
}

func (n Nap) Duration() int {
	return n.End.Minute() - n.Start.Minute()
}

type Shift struct {
	Guard int
	Start time.Time
	Naps  []Nap
}

func findShifts(events []*Event) []Shift {
	shifts := []Shift{}
	shift := Shift{Guard: -1}
	nap := Nap{}
	for _, event := range events {
		switch event.Type {
		case TYPE_BEGIN:
			// Special test for first shift.
			if shift.Guard != -1 {
				shifts = append(shifts, shift)
			}
			shift = Shift{Guard: event.Guard, Start: event.Time}
		case TYPE_SLEEP:
			nap = Nap{Start: event.Time}
		case TYPE_WAKE:
			if !nap.Start.IsZero() {
				nap.End = event.Time
				shift.Naps = append(shift.Naps, nap)
			}
		}
	}
	shifts = append(shifts, shift)
	return shifts
}

func sleepiestGuard(shifts []Shift) int {
	guards := make(map[int]int)
	for _, shift := range shifts {
		for _, nap := range shift.Naps {
			guards[shift.Guard] = guards[shift.Guard] + nap.Duration()
		}
	}

	sleepiest := 0
	for id, minutes := range guards {
		if guards[sleepiest] < minutes {
			sleepiest = id
		}
	}
	return sleepiest
}

func (s Shift) MinutesAsleep() int {
	minutes := 0
	for _, nap := range s.Naps {
		minutes = minutes + nap.Duration()
	}
	return minutes
}

func sleepiestMinute(naps []Nap) int {
	minutes := make([]int, 60)
	for _, nap := range naps {
		for i := nap.Start.Minute(); i < nap.End.Minute(); i++ {
			minutes[i] = minutes[i] + 1
		}
	}

	max_minute := 0
	for minute, naps := range minutes {
		if naps > minutes[max_minute] {
			max_minute = minute
		}
	}

	return max_minute
}

func sleepsPerMinute(shifts []Shift) map[int][]int {
	guards := make(map[int][]int)
	for _, shift := range shifts {
		if _, ok := guards[shift.Guard]; !ok {
			guards[shift.Guard] = make([]int, 60)
		}
		for _, nap := range shift.Naps {
			for i := nap.Start.Minute(); i < nap.End.Minute(); i++ {
				guards[shift.Guard][i] = guards[shift.Guard][i] + 1
			}
		}
	}
	return guards
}

func sleepiestMinuteAllTime(guards map[int][]int) (int, int, int) {
	sleepiest_guard := -1
	sleepiest_minute := -1
	sleepiest_count := -1
	for guard, minutes := range guards {
		for minute, count := range minutes {
			if count > sleepiest_count {
				sleepiest_guard = guard
				sleepiest_minute = minute
				sleepiest_count = count
			}
		}
	}
	return sleepiest_guard, sleepiest_minute, sleepiest_count
}
