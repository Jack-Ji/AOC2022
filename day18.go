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
	airs             [][][]bool
	xLen, yLen, zLen int
)

func getSurface1(cube Cube) int {
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

func getSurface2(cube Cube) int {
	var s = 0
	if cube.x == 0 || airs[cube.z][cube.y][cube.x-1] {
		s++
	}
	if cube.x == xLen-1 || airs[cube.z][cube.y][cube.x+1] {
		s++
	}
	if cube.y == 0 || airs[cube.z][cube.y-1][cube.x] {
		s++
	}
	if cube.y == yLen-1 || airs[cube.z][cube.y+1][cube.x] {
		s++
	}
	if cube.z == 0 || airs[cube.z-1][cube.y][cube.x] {
		s++
	}
	if cube.z == zLen-1 || airs[cube.z+1][cube.y][cube.x] {
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

	space = make([][][]bool, zLen)
	airs = make([][][]bool, zLen)
	for i := 0; i < zLen; i++ {
		space[i] = make([][]bool, yLen)
		airs[i] = make([][]bool, yLen)
		for j := 0; j < yLen; j++ {
			space[i][j] = make([]bool, xLen)
			airs[i][j] = make([]bool, xLen)
		}
	}
	for _, c := range cubes {
		space[c.z][c.y][c.x] = true
	}

	//part 1
	var surface = 0
	for _, c := range cubes {
		surface += getSurface1(c)
	}
	fmt.Println(surface)

	//part 2
	// expand air, from outside to inside
	var airN = 0
	for z := 0; z < zLen; z++ {
		for y := 0; y < yLen; y++ {
			for x := 0; x < xLen; x++ {
				if !(z == 0 || z == zLen-1 ||
					y == 0 || y == yLen-1 ||
					x == 0 || x == xLen-1) {
					continue
				}
				airs[z][y][x] = !space[z][y][x]
				if airs[z][y][x] {
					airN++
				}
			}
		}
	}
	for {
		var newAirN = 0
		for z := 0; z < zLen; z++ {
			for y := 0; y < yLen; y++ {
				for x := 0; x < xLen; x++ {
					if airs[z][y][x] || space[z][y][x] {
						continue
					}
					if airs[z][y][x-1] ||
						airs[z][y][x+1] ||
						airs[z][y-1][x] ||
						airs[z][y+1][x] ||
						airs[z-1][y][x] ||
						airs[z+1][y][x] {
						airs[z][y][x] = true
						newAirN++
						continue
					}
				}
			}
		}
		if newAirN == 0 {
			break
		}
		airN += newAirN
	}
	surface = 0
	for _, c := range cubes {
		surface += getSurface2(c)
	}
	fmt.Println(surface)
}
