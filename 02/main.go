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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		two, three := count(scanner.Text())
		if two {
			twice += 1
		}
		if three {
			thrice += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	log.Println("A:", twice*thrice)
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

func differsByOne(a, b string) bool {
	return false
}
