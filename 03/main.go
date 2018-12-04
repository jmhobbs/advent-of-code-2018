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

type Claim struct {
	ID int
	X  int
	Y  int
	W  int
	H  int
}

func ParseClaim(s string) Claim {
	return Claim{}
}

type Fabric struct {
	Size int
}

func NewFabric(size int) *Fabric {
	return &Fabric{size}
}

func (f *Fabric) AddClaim(c Claim) {
}

func (f *Fabric) OverlappingInches() int {
	return 0
}
