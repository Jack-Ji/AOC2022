package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Dir int

const (
	right Dir = iota
	down
	left
	up
)

func (d Dir) turn(b byte) Dir {
	var v = int(d)
	switch b {
	case 'R':
		v = (v + 1) % 4
	case 'L':
		v -= 1
		if v < 0 {
			v = 3
		}
	default:
		panic("unreachable")
	}
	return Dir(v)
}

func (d Dir) move(steps int) {
	switch d {
	case right:
		for steps > 0 {
			if tilemap[py][px] != none {
				steps--
			}
			if tilemap[py][px] != wall {
				px = (px + 1) % width
			}
		}
		for tilemap[py][px] != floor {
			px = ((px - 1) + width) % width
		}
	case down:
		for steps > 0 {
			if tilemap[py][px] != none {
				steps--
			}
			if tilemap[py][px] != wall {
				py = (py + 1) % height
			}
		}
		for tilemap[py][px] != floor {
			py = ((py - 1) + height) % height
		}
	case left:
		for steps > 0 {
			if tilemap[py][px] != none {
				steps--
			}
			if tilemap[py][px] != wall {
				px = ((px - 1) + width) % width
			}
		}
		for tilemap[py][px] != floor {
			px = (px + 1) % width
		}
	case up:
		for steps > 0 {
			if tilemap[py][px] != none {
				steps--
			}
			if tilemap[py][px] != wall {
				py = ((py - 1) + height) % height
			}
		}
		for tilemap[py][px] != floor {
			py = (py + 1) % height
		}
	}
}

type Inst struct {
	steps int
	dir   Dir
}

func parseInstructions(text string) {
	var (
		inst Inst
		off  = 0
		idx  = 0
	)

	for {
		if text[idx] >= '0' && text[idx] <= '9' {
			idx++
		} else {
			inst.steps, _ = strconv.Atoi(text[off:idx])
			instructions = append(instructions, inst)
			inst.dir = inst.dir.turn(text[idx])

			idx++
			off = idx
		}

		if idx == len(text) {
			inst.steps, _ = strconv.Atoi(text[off:idx])
			instructions = append(instructions, inst)
			break
		}
	}
}

type Tile int

const (
	none Tile = iota
	floor
	wall
)

var (
	instructions  []Inst
	tilemap       [][]Tile
	width, height int
	px, py        int
)

func printMap() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == py && j == px {
				fmt.Printf("@")
				continue
			}
			switch tilemap[i][j] {
			case none:
				fmt.Printf("~")
			case floor:
				fmt.Printf(".")
			case wall:
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	bs, _ := os.ReadFile("day22.txt")
	lines := strings.Split(strings.TrimRightFunc(string(bs),
		unicode.IsSpace), "\n")

	var readInst bool
	for i, s := range lines {
		l := strings.TrimRightFunc(s, unicode.IsSpace)
		if readInst {
			parseInstructions(l)
			break
		}
		if l == "" {
			height = i
			readInst = true
		} else if len(l) > width {
			width = len(l)
		}
	}
	tilemap = make([][]Tile, height)
	for i := 0; i < height; i++ {
		tilemap[i] = make([]Tile, width)
	}
	for i, s := range lines {
		l := strings.TrimRightFunc(s, unicode.IsSpace)
		if l == "" {
			break
		}
		for j, c := range l {
			if c == ' ' {
				continue
			}
			if i == 0 && px == 0 {
				px = j
			}
			if c == '.' {
				tilemap[i][j] = floor
			} else {
				tilemap[i][j] = wall
			}
		}
	}

	//part 1
	for _, inst := range instructions {
		inst.dir.move(inst.steps)
	}
	fmt.Println(
		px+1,
		py+1,
		(py+1)*1000+(px+1)*4+int(instructions[len(instructions)-1].dir),
	)
}
