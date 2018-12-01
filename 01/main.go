package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var shifts []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shifts = append(shifts, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	log.Println("A:", frequencyShift(shifts))
	log.Println("B:", duplicateFrequencyShift(shifts))
}

func duplicateFrequencyShift(shifts []string) int {
	v := 0
	seen := map[int]bool{0: true}
	for {
		for _, shift := range shifts {
			v = v + parseShift(shift)
			if _, ok := seen[v]; ok {
				return v
			}
			seen[v] = true
		}
	}
}

func frequencyShift(shifts []string) int {
	v := 0
	for _, shift := range shifts {
		v = v + parseShift(shift)
	}
	return v
}

func parseShift(shift string) int {
	sign := shift[0]
	v, err := strconv.Atoi(shift[1:])
	if err != nil {
		panic(err)
	}
	if sign == '-' {
		v = v * -1
	}
	return v
}
