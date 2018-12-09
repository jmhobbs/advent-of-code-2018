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

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}
}

func reduce(input string) string {
	var last rune
	var output []rune
	var reduced bool
	for {
		reduced = false
		last = 0
		output = output[:0]

		for _, c := range input {
			if c-32 == last || c+32 == last {
				last = 0
				reduced = true
				continue
			}
			if last != 0 {
				output = append(output, last)
			}
			last = c
		}
		if last != 0 {
			output = append(output, last)
		}

		if !reduced {
			break
		}
		input = string(output)
	}
	return input
}
