package main

import (
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
