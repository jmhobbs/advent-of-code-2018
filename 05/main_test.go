package main

import (
	"testing"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

func TestReduce(t *testing.T) {

	/*
		In aA, a and A react, leaving nothing behind.
		In abBA, bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing.
		In abAB, no two adjacent units are of the same type, and so nothing happens.
		In aabAAB, even though aa and AA are of the same type, their polarities match, and so nothing happens.
	*/
	samples := []struct {
		Input  string
		Output string
	}{
		{
			"aA",
			"",
		},
		{
			"abBA",
			"",
		},
		{
			"abAB",
			"abAB",
		},
		{
			"aabAAB",
			"aabAAB",
		},
		/*
			Now, consider a larger example, dabAcCaCBAcCcaDA:

			dabAcCaCBAcCcaDA  The first 'cC' is removed.
			dabAaCBAcCcaDA    This creates 'Aa', which is removed.
			dabCBAcCcaDA      Either 'cC' or 'Cc' are removed (the result is the same).
			dabCBAcaDA        No further actions can be taken.
		*/
		{
			"dabAcCaCBAcCcaDA",
			"dabCBAcaDA",
		},
	}

	for _, sample := range samples {
		h.Equals(t, sample.Output, reduce(sample.Input))
	}
}
