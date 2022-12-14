package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BlockType int

const (
	air BlockType = iota
	rock
	sand
)

var (
	minX   = 1000
	maxX   = 0
	maxY   = 0
	width  int
	height int
	dropX  = 500
	dropY  = 0
	blocks []BlockType
)

func posToIndex(x, y int) int {
	return y*width + (x - minX)
}

func printMap() {
	fmt.Println("====================================")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			switch blocks[i*width+j] {
			case air:
				fmt.Printf(".")
			case rock:
				fmt.Printf("#")
			case sand:
				fmt.Printf("o")
			}
		}
		fmt.Println()
	}
}

// Simulate dropping a sand,returns false if sand falls through
// walls or can't fall at all
func simulate() bool {
	var (
		x = dropX
		y = dropY
	)

	for y < height-1 {
		// Test down
		idx := posToIndex(x, y+1)
		if blocks[idx] == air {
			y += 1
			continue
		}

		// Test down-left
		idx = posToIndex(x-1, y+1)
		if blocks[idx] == air {
			x -= 1
			y += 1
			continue
		}

		// Test down-right
		idx = posToIndex(x+1, y+1)
		if blocks[idx] == air {
			x += 1
			y += 1
			continue
		}

		// Rest in place
		blocks[posToIndex(x, y)] = sand
		return y != dropY
	}

	return false
}

func main() {
	bs, _ := os.ReadFile("day14.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	// Initialize map
	for _, line := range lines {
		coords := strings.Split(strings.TrimSpace(line), "->")
		for _, coord := range coords {
			c := strings.TrimSpace(coord)
			idx := strings.IndexByte(c, ',')
			x, _ := strconv.Atoi(c[:idx])
			y, _ := strconv.Atoi(c[idx+1:])
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	minX -= 200
	maxX += 200
	maxY += 3
	width = maxX - minX + 1
	height = maxY
	blocks = make([]BlockType, width*height)
	for x := minX; x <= maxX; x++ {
		blocks[posToIndex(x, height-1)] = rock
	}
	for _, line := range lines {
		coords := strings.Split(strings.TrimSpace(line), "->")
		if len(coords) < 2 {
			panic("unreachable")
		}

		for i := 0; i < len(coords)-1; i++ {
			c := strings.TrimSpace(coords[i])
			idx := strings.IndexByte(c, ',')
			x0, _ := strconv.Atoi(c[:idx])
			y0, _ := strconv.Atoi(c[idx+1:])

			c = strings.TrimSpace(coords[i+1])
			idx = strings.IndexByte(c, ',')
			x1, _ := strconv.Atoi(c[:idx])
			y1, _ := strconv.Atoi(c[idx+1:])

			if x0 == x1 {
				var step = 1
				if y0 > y1 {
					step = -1
				}
				for y := y0; y != y1; y += step {
					blocks[posToIndex(x0, y)] = rock
				}
				blocks[posToIndex(x1, y1)] = rock
			} else if y0 == y1 {
				var step = 1
				if x0 > x1 {
					step = -1
				}
				for x := x0; x != x1; x += step {
					blocks[posToIndex(x, y0)] = rock
				}
				blocks[posToIndex(x1, y1)] = rock
			} else {
				panic("unreachable")
			}
		}
	}

	var sandNum = 0
	for simulate() {
		sandNum++
	}
	printMap()
	fmt.Println(sandNum)
}
