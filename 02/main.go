package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func count(s string) (bool, bool) {
	v := make(map[rune]uint)
	for _, c := range s {
		v[c] = v[c] + 1
	}

	two := false
	three := false
	for _, i := range v {
		two = two || (i == 2)
		three = three || (i == 3)
	}

	return two, three
}
