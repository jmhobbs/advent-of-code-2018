package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	twice := 0
	thrice := 0
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		two, three := count(line)
		if two {
			twice += 1
		}
		if three {
			thrice += 1
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	log.Println("A:", twice*thrice)
	log.Println("B:", samsies(lines))
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

func differsByOneIndex(a, b string) int {
	differs := -1
	for i, r := range a {
		if b[i] != byte(r) {
			if differs != -1 {
				return -1
			}
			differs = i
		}
	}
	return differs
}

func samsies(lines []string) string {
	var idx int
	for i, outer := range lines {
		for _, inner := range lines[i:] {
			idx = differsByOneIndex(outer, inner)
			if idx != -1 {
				return outer[:idx] + outer[idx+1:]
			}
		}
	}
	panic("I should not have done that.")
}
