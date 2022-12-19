package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

var (
	cubes            []Cube
	space            [][][]bool
	xLen, yLen, zLen int
)

func getSurface(cube Cube) int {
	var s = 0
	if cube.x == 0 || !space[cube.z][cube.y][cube.x-1] {
		s++
	}
	if cube.x == xLen-1 || !space[cube.z][cube.y][cube.x+1] {
		s++
	}
	if cube.y == 0 || !space[cube.z][cube.y-1][cube.x] {
		s++
	}
	if cube.y == yLen-1 || !space[cube.z][cube.y+1][cube.x] {
		s++
	}
	if cube.z == 0 || !space[cube.z-1][cube.y][cube.x] {
		s++
	}
	if cube.z == zLen-1 || !space[cube.z+1][cube.y][cube.x] {
		s++
	}
	return s
}

func main() {
	bs, _ := os.ReadFile("day18.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	for _, s := range lines {
		l := strings.TrimSpace(s)
		pos := strings.Split(l, ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		z, _ := strconv.Atoi(pos[2])
		if x+1 > xLen {
			xLen = x + 1
		}
		if y+1 > yLen {
			yLen = y + 1
		}
		if z+1 > zLen {
			zLen = z + 1
		}
		cubes = append(cubes, Cube{x: x, y: y, z: z})
	}

	// Init space
	space = make([][][]bool, zLen)
	for i := 0; i < zLen; i++ {
		space[i] = make([][]bool, yLen)
		for j := 0; j < yLen; j++ {
			space[i][j] = make([]bool, xLen)
		}
	}
	for _, c := range cubes {
		space[c.z][c.y][c.x] = true
	}

	var surface = 0
	for _, c := range cubes {
		surface += getSurface(c)
	}
	fmt.Println(surface)
}
