package main

import (
	"fmt"
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

func (d Dir) next(x, y int) string {
	var nx, ny = x, y
	switch d {
	case N:
		ny--
		if ny == 0 {
			ny = height - 2
		}
	case S:
		ny++
		if ny == height-1 {
			ny = 1
		}
	case W:
		nx--
		if nx == 0 {
			nx = width - 2
		}
	case E:
		nx++
		if nx == width-1 {
			nx = 1
		}
	}
	return getKey(nx, ny)
}

type Blizzard struct {
	dir Dir
}

var (
	blizzards = map[string][]Blizzard{}
	width     int
	height    int
	dstx      int
	dsty      int
	px        int = 1
	py        int = 0
)

func getKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getXY(key string) (int, int) {
	var x, y int
	fmt.Sscanf(key, "%d,%d", &x, &y)
	return x, y
}

func isSuitable(x, y int, bds map[string][]Blizzard) bool {
	if x == 1 && y == 0 {
		return true
	}

	if x == dstx && y == dsty {
		return true
	}

	if x <= 0 || x >= width-1 || y <= 0 || y >= height-1 {
		return false
	}

	if bds[getKey(x, y)] != nil {
		return false
	}

	return true
}

var level = 0

func calcMinutes(x, y, min, threshold int, bds map[string][]Blizzard) int {
	level++
	defer func() {
		fmt.Println(level, ">", x, y, min, threshold)
		level--
	}()

	if x == dstx && y == dsty {
		return min
	}

	if min >= threshold-1 {
		return threshold
	}

	var (
		result = threshold
		nbs    = map[string][]Blizzard{}
	)
	for k, bs := range bds {
		for _, b := range bs {
			np := b.dir.next(getXY(k))
			if bs := nbs[np]; bs == nil {
				nbs[np] = []Blizzard{b}
			} else {
				bs = append(bs, b)
				nbs[np] = bs
			}
		}
	}

	if isSuitable(x-1, y, nbs) {
		nmin := calcMinutes(x-1, y, min+1, threshold, nbs)
		if result > nmin {
			result = nmin
		}
		if result < threshold {
			threshold = result
		}
	}

	if isSuitable(x+1, y, nbs) {
		nmin := calcMinutes(x+1, y, min+1, threshold, nbs)
		if result > nmin {
			result = nmin
		}
		if result < threshold {
			threshold = result
		}
	}

	if isSuitable(x, y-1, nbs) {
		nmin := calcMinutes(x, y-1, min+1, threshold, nbs)
		if result > nmin {
			result = nmin
		}
		if result < threshold {
			threshold = result
		}
	}

	if isSuitable(x, y+1, nbs) {
		nmin := calcMinutes(x, y+1, min+1, threshold, nbs)
		if result > nmin {
			result = nmin
		}
		if result < threshold {
			threshold = result
		}
	}

	if isSuitable(x, y, nbs) {
		nmin := calcMinutes(x, y, min+1, threshold, nbs)
		if result > nmin {
			result = nmin
		}
		if result < threshold {
			threshold = result
		}
	}

	return result
}

func printMap(bds map[string][]Blizzard) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if j == px && i == py {
				fmt.Printf("E")
			} else if i == 0 && j == 1 {
				fmt.Printf(".")
			} else if i == dsty && j == dstx {
				fmt.Printf(".")
			} else if j == 0 || j == width-1 || i == 0 || i == height-1 {
				fmt.Printf("#")
			} else if bs := bds[getKey(j, i)]; bs != nil {
				if len(bs) > 1 {
					fmt.Printf("%d", len(bs))
				} else {
					switch bs[0].dir {
					case N:
						fmt.Printf("^")
					case S:
						fmt.Printf("v")
					case W:
						fmt.Printf("<")
					case E:
						fmt.Printf(">")
					}
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
}

func main() {
	bs, _ := os.ReadFile("day24.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	height = len(lines)
	for i, s := range lines {
		l := strings.TrimSpace(s)
		width = len(l)
		for j, c := range l {
			switch c {
			case '^':
				blizzards[getKey(j, i)] = []Blizzard{Blizzard{dir: N}}
			case '<':
				blizzards[getKey(j, i)] = []Blizzard{Blizzard{dir: W}}
			case '>':
				blizzards[getKey(j, i)] = []Blizzard{Blizzard{dir: E}}
			case 'v':
				blizzards[getKey(j, i)] = []Blizzard{Blizzard{dir: S}}
			default:
				continue
			}
		}
	}
	dstx = width - 2
	dsty = height - 1

	// part1
	fmt.Println(calcMinutes(px, py, 0, 360, blizzards))
}
