package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Dir int

const (
	N Dir = iota
	S
	W
	E
)

var dir Dir

func (d Dir) move() Dir {
	var i = int(d)
	i = (i + 1) % 4
	return Dir(i)
}

type Elve struct {
}

func (e Elve) getNextPosition(x, y int) (string, bool) {
	var (
		d = dir
	)

	if elves[getKey(x-1, y-1)] == nil &&
		elves[getKey(x, y-1)] == nil &&
		elves[getKey(x+1, y-1)] == nil &&
		elves[getKey(x+1, y)] == nil &&
		elves[getKey(x+1, y+1)] == nil &&
		elves[getKey(x, y+1)] == nil &&
		elves[getKey(x-1, y+1)] == nil &&
		elves[getKey(x-1, y)] == nil {
		return "", false
	}

	for i := 0; i < 4; i++ {
		switch d {
		case N:
			if elves[getKey(x-1, y-1)] == nil &&
				elves[getKey(x, y-1)] == nil &&
				elves[getKey(x+1, y-1)] == nil {
				return getKey(x, y-1), true
			}
		case S:
			if elves[getKey(x-1, y+1)] == nil &&
				elves[getKey(x, y+1)] == nil &&
				elves[getKey(x+1, y+1)] == nil {
				return getKey(x, y+1), true
			}
		case W:
			if elves[getKey(x-1, y-1)] == nil &&
				elves[getKey(x-1, y)] == nil &&
				elves[getKey(x-1, y+1)] == nil {
				return getKey(x-1, y), true
			}
		case E:
			if elves[getKey(x+1, y-1)] == nil &&
				elves[getKey(x+1, y)] == nil &&
				elves[getKey(x+1, y+1)] == nil {
				return getKey(x+1, y), true
			}
		}
		d = d.next()
	}
	return "", false
}

var elves = map[string]*Elve{}

func getKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getXY(key string) (int, int) {
	var x, y int
	fmt.Sscanf(key, "%d,%d", &x, &y)
	return x, y
}

func countEmptyTiles() int {
	var (
		minx int = math.MaxInt
		miny int = math.MaxInt
		maxx int = 0
		maxy int = 0
	)

	for k := range elves {
		x, y := getXY(k)
		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
	}

	var tilecount int
	for x := minx; x <= maxx; x++ {
		for y := miny; y <= maxy; y++ {
			key := getKey(x, y)
			if elves[key] == nil {
				tilecount++
			}
		}
	}
	return tilecount
}

func printMap() {
	var (
		minx int = math.MaxInt
		miny int = math.MaxInt
		maxx int = 0
		maxy int = 0
	)

	for k := range elves {
		x, y := getXY(k)
		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
	}

	fmt.Printf("X:[%d-%d]  Y:[%d-%d]\n", minx, maxx, miny, maxy)
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			key := getKey(x, y)
			if elves[key] == nil {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n-----------------------------------------------")
}

func main() {
	bs, _ := os.ReadFile("day23.txt")
	lines := strings.Split(string(bs), "\n")

	for i, s := range lines {
		for j, c := range strings.TrimSpace(s) {
			if c == '#' {
				elves[getKey(j, i)] = &Elve{}
			}
		}
	}

	// part1
	//round := 10
	//for round > 0 {
	//	var moves = map[string]string{} // to->from
	//	for k, v := range elves {
	//		to, moved := v.getNextPosition(getXY(k))
	//		if moved {
	//			if moves[to] == "" {
	//				moves[to] = k
	//			} else {
	//				delete(moves, to)
	//			}
	//		}
	//	}
	//	if len(moves) == 0 {
	//		break
	//	} else {
	//		round--
	//		dir = dir.next()

	//		for to, from := range moves {
	//			e := elves[from]
	//			delete(elves, from)
	//			elves[to] = e
	//		}
	//	}
	//}
	//printMap()
	//fmt.Println(countEmptyTiles())

	// part2
	round := 0
	for {
		round++
		var moves = map[string]string{} // to->from
		for k, v := range elves {
			to, moved := v.getNextPosition(getXY(k))
			if moved {
				if moves[to] == "" {
					moves[to] = k
				} else {
					delete(moves, to)
				}
			}
		}
		if len(moves) == 0 {
			break
		} else {
			dir = dir.next()

			for to, from := range moves {
				e := elves[from]
				delete(elves, from)
				elves[to] = e
			}
		}
	}
	printMap()
	fmt.Println(round)
}
