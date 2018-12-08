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
		h.Equals(t, expect.Time, evt.Time.Format(EVENT_TIME_FORMAT))
		h.Equals(t, expect.Type, evt.Type)
		h.Equals(t, expect.Guard, evt.Guard)
	}
}

func TestSorting(t *testing.T) {
	lines := []string{
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:05] falls asleep",
	}

	events := []*Event{}
	for _, line := range lines {
		events = append(events, parseLine(line))
	}

	sorted := sortEvents(events)

	h.Equals(t, "1518-11-01 00:00", sorted[0].Time.Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:05", sorted[1].Time.Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:25", sorted[2].Time.Format(EVENT_TIME_FORMAT))
	h.Equals(t, "1518-11-01 00:30", sorted[3].Time.Format(EVENT_TIME_FORMAT))
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
			time.Date(1518, time.November, 1, 0, 25, 0, 0, time.UTC),
		},
		Nap{
			time.Date(1518, time.November, 3, 0, 24, 0, 0, time.UTC),
			time.Date(1518, time.November, 3, 0, 29, 0, 0, time.UTC),
		},
		Nap{
			time.Date(1518, time.November, 1, 0, 30, 0, 0, time.UTC),
			time.Date(1518, time.November, 1, 0, 55, 0, 0, time.UTC),
		},
	}

	h.Equals(t, 24, sleepiestMinute(naps))
}

func TestFindShifts(t *testing.T) {
	events := []*Event{
		&Event{
			time.Date(1518, time.November, 1, 0, 0, 0, 0, time.UTC),
			10,
			TYPE_BEGIN,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC),
			-1,
			TYPE_SLEEP,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 25, 0, 0, time.UTC),
			-1,
			TYPE_WAKE,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 30, 0, 0, time.UTC),
			8,
			TYPE_BEGIN,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 35, 0, 0, time.UTC),
			-1,
			TYPE_SLEEP,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 37, 0, 0, time.UTC),
			-1,
			TYPE_WAKE,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 42, 0, 0, time.UTC),
			-1,
			TYPE_SLEEP,
		},
		&Event{
			time.Date(1518, time.November, 1, 0, 55, 0, 0, time.UTC),
			-1,
			TYPE_WAKE,
		},
	}

	expected := []Shift{
		Shift{
			10,
			time.Date(1518, time.November, 1, 0, 0, 0, 0, time.UTC),
			[]Nap{
				Nap{
					time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 25, 0, 0, time.UTC),
				},
			},
		},
		Shift{
			8,
			time.Date(1518, time.November, 1, 0, 30, 0, 0, time.UTC),
			[]Nap{
				Nap{
					time.Date(1518, time.November, 1, 0, 35, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 37, 0, 0, time.UTC),
				},
				Nap{
					time.Date(1518, time.November, 1, 0, 42, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 55, 0, 0, time.UTC),
				},
			},
		},
	}

	h.Equals(t, expected, findShifts(events))
}

func TestSleepiestGuard(t *testing.T) {
	shifts := []Shift{
		Shift{
			10,
			time.Date(1518, time.November, 1, 0, 0, 0, 0, time.UTC),
			[]Nap{
				Nap{
					time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 25, 0, 0, time.UTC),
				},
			},
		},
		Shift{
			8,
			time.Date(1518, time.November, 1, 0, 30, 0, 0, time.UTC),
			[]Nap{
				Nap{
					time.Date(1518, time.November, 1, 0, 35, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 37, 0, 0, time.UTC),
				},
				Nap{
					time.Date(1518, time.November, 1, 0, 42, 0, 0, time.UTC),
					time.Date(1518, time.November, 1, 0, 55, 0, 0, time.UTC),
				},
			},
		},
	}

	h.Equals(t, 10, sleepiestGuard(shifts))
}

func TestMinutesAsleep(t *testing.T) {
	s := Shift{
		0,
		time.Date(1518, time.November, 1, 0, 0, 0, 0, time.UTC),
		[]Nap{
			Nap{
				time.Date(1518, time.November, 1, 0, 5, 0, 0, time.UTC),
				time.Date(1518, time.November, 1, 0, 15, 0, 0, time.UTC),
			},
			Nap{
				time.Date(1518, time.November, 1, 0, 20, 0, 0, time.UTC),
				time.Date(1518, time.November, 1, 0, 23, 0, 0, time.UTC),
			},
		},
	}

	h.Equals(t, 13, s.MinutesAsleep())
}
