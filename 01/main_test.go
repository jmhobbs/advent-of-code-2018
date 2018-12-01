package main

import (
	"testing"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

type test struct {
	Shifts   []string
	Expected int
}

func TestShifts(t *testing.T) {
	tests := []test{
		/*
			Current frequency  0, change of +1; resulting frequency  1.
			Current frequency  1, change of -2; resulting frequency -1.
			Current frequency -1, change of +3; resulting frequency  2.
			Current frequency  2, change of +1; resulting frequency  3.
		*/
		{
			[]string{"+1"},
			1,
		},
		{
			[]string{"+1", "-2"},
			-1,
		},
		{
			[]string{"+1", "-2", "+3"},
			2,
		},
		{
			[]string{"+1", "-2", "+3", "+1"},
			3,
		},
		// +1, +1, +1 results in  3
		{
			[]string{"+1", "+1", "-2"},
			0,
		},
		// +1, +1, -2 results in  0
		{
			[]string{"+1", "+1", "-2"},
			0,
		},
		// -1, -2, -3 results in -6
		{
			[]string{"-1", "-2", "-3"},
			-6,
		},
	}

	var actual int
	for _, test := range tests {
		actual = frequencyShift(test.Shifts)
		h.Equals(t, test.Expected, actual)
	}
}

func TestParseShift(t *testing.T) {
	h.Equals(t, parseShift("+5"), 5)
	h.Equals(t, parseShift("-5"), -5)
}

func TestDuplicateShifts(t *testing.T) {
	/*
		Current frequency  0, change of +1; resulting frequency  1.
		Current frequency  1, change of -2; resulting frequency -1.
		Current frequency -1, change of +3; resulting frequency  2.
		Current frequency  2, change of +1; resulting frequency  3.
		(At this point, the device continues from the start of the list.)
		Current frequency  3, change of +1; resulting frequency  4.
		Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
	*/
	tests := []test{
		{
			[]string{"+1", "-2", "+3", "+1"},
			2,
		},
		// +1, -1 first reaches 0 twice.
		{
			[]string{"+1", "-1"},
			0,
		},
		// +3, +3, +4, -2, -4 first reaches 10 twice.
		{
			[]string{"+3", "+3", "+4", "-2", "-4"},
			10,
		},
		// -6, +3, +8, +5, -6 first reaches 5 twice.
		{
			[]string{"-6", "+3", "+8", "+5", "-6"},
			5,
		},
		// +7, +7, -2, -7, -4 first reaches 14 twice.
		{
			[]string{"+7", "+7", "-2", "-7", "-4"},
			14,
		},
	}

	var actual int
	for _, test := range tests {
		actual = duplicateFrequencyShift(test.Shifts)
		h.Equals(t, test.Expected, actual)
	}
}
