package main

import (
	"testing"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

type testCase struct {
	String     string
	TwoCount   bool
	ThreeCount bool
}

var testCases []testCase = []testCase{
	// abcdef contains no letters that appear exactly two or three times.
	{
		"abcdef",
		false,
		false,
	},
	// bababc contains two a and three b, so it counts for both.
	{
		"bababc",
		true,
		true,
	},
	// abbcde contains two b, but no letter appears exactly three times.
	{
		"abbcde",
		true,
		false,
	},
	// abcccd contains three c, but no letter appears exactly two times.
	{
		"abcccd",
		false,
		true,
	},
	// aabcdd contains two a and two d, but it only counts once.
	{
		"aabcdd",
		true,
		false,
	},
	// abcdee contains two e.
	{
		"abcdee",
		true,
		false,
	},
	// ababab contains three a and three b, but it only counts once.
	{
		"ababab",
		false,
		true,
	},
}

func TestCount(t *testing.T) {
	for _, c := range testCases {
		two, three := count(c.String)
		h.Equals(t, c.TwoCount, two)
		h.Equals(t, c.ThreeCount, three)
	}
}

func TestDiffersByOne(t *testing.T) {
	h.Equals(t, differsByOneIndex("abcde", "abcde"), -1)
	h.Equals(t, differsByOneIndex("abcde", "axcye"), -1)
	h.Equals(t, differsByOneIndex("fghij", "fguij"), 2)
}

func TestSamsies(t *testing.T) {
	lines := []string{
		"opitlop",
		"abcoghj",
		"nopenop",
		"abckghj",
		"wahnaha",
	}

	h.Equals(t, samsies(lines), "abcghj")
}
