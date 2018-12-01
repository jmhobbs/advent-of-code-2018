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
			[]string{"+1", "-2", "+3"},
			2,
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
