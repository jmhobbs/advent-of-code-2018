package main

import (
	"testing"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

func TestParseClaim(t *testing.T) {
	claim := ParseClaim("#123 @ 3,2: 5x4")
	h.Equals(t, 123, claim.ID)
	h.Equals(t, 3, claim.X)
	h.Equals(t, 2, claim.Y)
	h.Equals(t, 5, claim.W)
	h.Equals(t, 4, claim.H)
}

func TestOverlappingClaims(t *testing.T) {
	/*
		#1 @ 1,3: 4x4
		#2 @ 3,1: 4x4
		#3 @ 5,5: 2x2
	*/
	f := NewFabric(8)
	f.AddClaim(ParseClaim("#1 @ 1,3: 4x4"))
	f.AddClaim(ParseClaim("#2 @ 3,1: 4x4"))
	f.AddClaim(ParseClaim("#3 @ 5,5: 2x2"))
	h.Equals(t, f.OverlappingInches(), 4)
}
