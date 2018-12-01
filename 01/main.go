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

func frequencyShift(shifts []string) int {
	return 0
}
