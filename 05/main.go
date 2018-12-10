package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("A:", len(reduce(input)))
}

func reduce(input []byte) []byte {
	for {
		length := len(input)
		input = reduceStep(input)
		if len(input) == length {
			break
		}
	}

	return input
}

func reduceStep(input []byte) []byte {
	var last byte = 0
	for i, c := range input {
		if c-32 == last || c+32 == last {
			input = append(input[:i-1], input[i+1:]...)
			break
		}
		last = c
	}
	return input
}

func stripPolymer(input []byte, polymer byte) []byte {
	return bytes.Replace(bytes.Replace(input, []byte{polymer}, []byte{}, -1), []byte{polymer + 32}, []byte{}, -1)
}

func findAllPolymers(input []byte) []byte {
	polymers := []byte{}
	for _, c := range input {
		if c >= 'A' && c <= 'Z' {
			if !bytes.ContainsRune(polymers, rune(c)) {
				polymers = append(polymers, c)
			}
		} else if c >= 'a' && c <= 'z' {
			if !bytes.ContainsRune(polymers, rune(c-32)) {
				polymers = append(polymers, c-32)
			}
		}
	}
	return polymers
}

// Wrote this first, it does not work how the description does.
func wrongReduce(input string) string {
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
