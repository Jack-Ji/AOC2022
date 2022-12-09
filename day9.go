package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x int
	y int
}

var (
	knots = []Knot{
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
		Knot{100, 100},
	}
	visited = map[string]bool{
		"100,100": true,
	}
)

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

func move(h, t *Knot) bool {
	var (
		moveX int
		moveY int
	)

	if abs(h.y-t.y) > 1 || abs(h.x-t.x) > 1 {
		if h.x > t.x {
			moveX = 1
		} else if h.x < t.x {
			moveX = -1
		}

		if h.y > t.y {
			moveY = 1
		} else if h.y < t.y {
			moveY = -1
		}
	}

	if moveX == 0 && moveY == 0 {
		return false
	}
	t.x += moveX
	t.y += moveY
	return true
}

func step(c byte) {
	switch c {
	case 'R':
		knots[0].x += 1
	case 'U':
		knots[0].y -= 1
	case 'L':
		knots[0].x -= 1
	case 'D':
		knots[0].y += 1
	default:
		panic("unreachable")
	}

	for i := 0; i < len(knots)-1; i++ {
		head := &knots[i]
		tail := &knots[i+1]
		if !move(head, tail) {
			return
		}
	}

	visited[fmt.Sprintf("%d,%d", knots[9].x, knots[9].y)] = true
}

func stepn(c byte, count int) {
	for count > 0 {
		step(c)
		count--
	}

	/*
		for j := 90; j < 110; j++ {
		LOOP:
			for i := 85; i < 115; i++ {
				for k, n := range knots {
					if n.x == i && n.y == j {
						fmt.Printf("%d", k)
						continue LOOP
					}
				}
				if i == 100 && j == 100 {
					fmt.Printf("s")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println("==================================")
	*/
}

func main() {
	bs, _ := os.ReadFile("day9.txt")
	fs := strings.Split(string(bs), "\n")

	for _, f := range fs {
		line := strings.TrimSpace(f)
		if len(line) == 0 {
			break
		}

		count, _ := strconv.Atoi(line[2:])
		stepn(line[0], count)
	}
	fmt.Println(len(visited))
}
