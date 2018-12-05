package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fabric := NewFabric(1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fabric.AddClaim(ParseClaim(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	log.Println("A:", fabric.OverlappingInches())
	log.Println("B:", fabric.NonOverlappingClaims()[0])
}

type Claim struct {
	ID int
	X  int
	Y  int
	W  int
	H  int
}

var claimRegexp *regexp.Regexp

func init() {
	claimRegexp = regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
}

func ParseClaim(s string) Claim {
	matches := claimRegexp.FindStringSubmatch(s)

	id, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(matches[4])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(matches[5])
	if err != nil {
		panic(err)
	}

	return Claim{
		ID: id,
		X:  x,
		Y:  y,
		W:  w,
		H:  h,
	}
}

type Fabric struct {
	Size   int
	grid   [][][]int
	claims []Claim
}

func NewFabric(size int) *Fabric {
	grid := make([][][]int, size)
	for i := range grid {
		grid[i] = make([][]int, size)
	}
	return &Fabric{Size: size, grid: grid}
}

func (f *Fabric) AddClaim(c Claim) {
	for x := c.X; x < c.X+c.W; x++ {
		for y := c.Y; y < c.Y+c.H; y++ {
			f.grid[x][y] = append(f.grid[x][y], c.ID)
		}
	}
	f.claims = append(f.claims, c)
}

func (f *Fabric) OverlappingInches() int {
	overlaps := 0

	for x := 0; x < f.Size; x++ {
		for y := 0; y < f.Size; y++ {
			if len(f.grid[x][y]) > 1 {
				overlaps = overlaps + 1
			}
		}
	}

	return overlaps
}

func (f *Fabric) String() string {
	s := []string{}

	for y := 0; y < f.Size; y++ {
		row := []string{}
		for x := 0; x < f.Size; x++ {
			switch len(f.grid[x][y]) {
			case 0:
				row = append(row, ".")
			case 1:
				row = append(row, strconv.Itoa(f.grid[x][y][0]))
			default:
				row = append(row, "X")
			}
		}
		s = append(s, strings.Join(row, ""))
	}

	return strings.Join(s, "\n")
}

func (f *Fabric) NonOverlappingClaims() []int {
	nonOverlapping := []int{}
	for _, c := range f.claims {
		sentinel := false
		for x := c.X; x < c.X+c.W && !sentinel; x++ {
			for y := c.Y; y < c.Y+c.H && !sentinel; y++ {
				if len(f.grid[x][y]) > 1 {
					sentinel = true
				}
			}
		}
		if !sentinel {
			nonOverlapping = append(nonOverlapping, c.ID)
		}
	}
	return nonOverlapping
}
