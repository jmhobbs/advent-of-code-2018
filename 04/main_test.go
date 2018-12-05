package main

import (
	"testing"
	"time"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

func TestParsing(t *testing.T) {
	expected := []struct {
		Line  string
		Time  string
		Type  EventType
		Guard int
	}{
		{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"1518-11-01 00:00",
			TYPE_BEGIN,
			10,
		},
		{
			"[1518-11-01 00:05] falls asleep",
			"1518-11-01 00:05",
			TYPE_SLEEP,
			-1,
		},
		{
			"[1518-11-01 00:25] wakes up",
			"1518-11-01 00:25",
			TYPE_WAKE,
			-1,
		},
	}

	for _, expect := range expected {
		evt := parseLine(expect.Line)
		h.Equals(t, expect.Time, evt.Time().Format(EVENT_TIME_FORMAT))
		h.Equals(t, expect.Type, evt.Type())
		h.Equals(t, expect.Guard, evt.Guard())
	}
}

func TestSorting(t *testing.T) {
	lines := []string{
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:05] falls asleep",
	}

	events := []Event{}
	for _, line := range lines {
		events = append(events, parseLine(line))
	}

	sorted := sortEvents(events)

	h.Equals(t, "1518-11-01 00:00", sorted[0].Time().Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:05", sorted[1].Time().Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:25", sorted[2].Time().Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:30", sorted[3].Time().Format(EVENT_TIME_FORMAT))
}

func TestSleepiestMinute(t *testing.T) {
	/*
		            000000000011111111112222222222333333333344444444445555555555
		            012345678901234567890123456789012345678901234567890123456789
		11-01  #10  .....####################.....#########################.....
		11-03  #10  ........................#####...............................
	*/
	naps := []Nap{
		Nap{
			time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC),
			20,
		},
		Nap{
			time.Date(1518, time.November, 3, 2, 4, 0, 0, time.UTC),
			5,
		},
		Nap{
			time.Date(1518, time.November, 1, 3, 0, 0, 0, time.UTC),
			25,
		},
	}

	h.Equals(t, 24, sleepiestMinute(naps))
}
