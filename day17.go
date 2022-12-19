package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Rock int

const (
	stripeH Rock = iota
	plus
	lreverted
	stripeV
	cube
)

func (r Rock) getRows() []Line {
	switch r {
	case stripeH:
		return []Line{
			Line{cells: [7]byte{'.', '.', '@', '@', '@', '@', '.'}},
		}
	case plus:
		return []Line{
			Line{cells: [7]byte{'.', '.', '.', '@', '.', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '@', '@', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '.', '@', '.', '.', '.'}},
		}
	case lreverted:
		return []Line{
			Line{cells: [7]byte{'.', '.', '.', '.', '@', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '.', '.', '@', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '@', '@', '.', '.'}},
		}
	case stripeV:
		return []Line{
			Line{cells: [7]byte{'.', '.', '@', '.', '.', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '.', '.', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '.', '.', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '.', '.', '.', '.'}},
		}
	case cube:
		return []Line{
			Line{cells: [7]byte{'.', '.', '@', '@', '.', '.', '.'}},
			Line{cells: [7]byte{'.', '.', '@', '@', '.', '.', '.'}},
		}
	}
	panic("unreachable")
}

func (r Rock) getWidth() int {
	switch r {
	case stripeH:
		return 4
	case plus:
		return 3
	case lreverted:
		return 3
	case stripeV:
		return 1
	case cube:
		return 2
	}
	panic("unreachable")
}

var rockTypeIndex = 0

func getNextRock() Rock {
	var rock = Rock(rockTypeIndex)
	rockTypeIndex = (rockTypeIndex + 1) % 5
	return rock
}

type Wind int

const (
	left Wind = iota
	right
)

func (w Wind) String() string {
	if w == left {
		return "left"
	}
	return "right"
}

var windIndex = 0

func getNextWind() Wind {
	w := wind[windIndex]
	windIndex = (windIndex + 1) % len(wind)
	if w == '<' {
		return left
	} else {
		return right
	}
}

type Line struct {
	cells [7]byte
}

var (
	fallingSpace = []Line{}
	wind         string
	height       int
)

// Simulate a falling rock until it rest
func simulate(debug bool) {
	var (
		rock = getNextRock()
		x    = 2
		y    = height + 3
	)

	newRows := rock.getRows()
	fallingSpace = fallingSpace[:height]
	fallingSpace = append(fallingSpace, Line{cells: [7]byte{'.', '.', '.', '.', '.', '.', '.'}})
	fallingSpace = append(fallingSpace, Line{cells: [7]byte{'.', '.', '.', '.', '.', '.', '.'}})
	fallingSpace = append(fallingSpace, Line{cells: [7]byte{'.', '.', '.', '.', '.', '.', '.'}})
	for i := len(newRows) - 1; i >= 0; i-- {
		fallingSpace = append(fallingSpace, newRows[i])
	}

	if debug {
		for i := len(fallingSpace) - 1; i >= 0; i-- {
			fmt.Printf("% 4d|%s|\n", i+1, string(fallingSpace[i].cells[:]))
		}
		fmt.Println("----+-------+----")
	}

	defer func() {
		if debug {
			for i := len(fallingSpace) - 1; i >= 0; i-- {
				fmt.Printf("% 4d|%s|\n", i+1, string(fallingSpace[i].cells[:]))
			}
			fmt.Println("----+-------+----")
			fmt.Println()
		}
	}()

	for {
		w := getNextWind()
		switch w {
		case left:
			if x > 0 {
				switch rock {
				case stripeH:
					if fallingSpace[y].cells[x-1] == '.' {
						fallingSpace[y].cells[x-1] = '@'
						fallingSpace[y].cells[x+3] = '.'
						x--
					}
				case plus:
					if fallingSpace[y].cells[x] == '.' &&
						fallingSpace[y+1].cells[x-1] == '.' &&
						fallingSpace[y+2].cells[x] == '.' {
						fallingSpace[y].cells[x] = '@'
						fallingSpace[y].cells[x+1] = '.'
						fallingSpace[y+1].cells[x-1] = '@'
						fallingSpace[y+1].cells[x+2] = '.'
						fallingSpace[y+2].cells[x] = '@'
						fallingSpace[y+2].cells[x+1] = '.'
						x--
					}
				case lreverted:
					if fallingSpace[y].cells[x-1] == '.' &&
						fallingSpace[y+1].cells[x+1] == '.' &&
						fallingSpace[y+2].cells[x+1] == '.' {
						fallingSpace[y].cells[x-1] = '@'
						fallingSpace[y].cells[x+2] = '.'
						fallingSpace[y+1].cells[x+1] = '@'
						fallingSpace[y+1].cells[x+2] = '.'
						fallingSpace[y+2].cells[x+1] = '@'
						fallingSpace[y+2].cells[x+2] = '.'
						x--
					}
				case stripeV:
					if fallingSpace[y].cells[x-1] == '.' &&
						fallingSpace[y+1].cells[x-1] == '.' &&
						fallingSpace[y+2].cells[x-1] == '.' &&
						fallingSpace[y+3].cells[x-1] == '.' {
						fallingSpace[y].cells[x-1] = '@'
						fallingSpace[y].cells[x] = '.'
						fallingSpace[y+1].cells[x-1] = '@'
						fallingSpace[y+1].cells[x] = '.'
						fallingSpace[y+2].cells[x-1] = '@'
						fallingSpace[y+2].cells[x] = '.'
						fallingSpace[y+3].cells[x-1] = '@'
						fallingSpace[y+3].cells[x] = '.'
						x--
					}
				case cube:
					if fallingSpace[y].cells[x-1] == '.' &&
						fallingSpace[y+1].cells[x-1] == '.' {
						fallingSpace[y].cells[x-1] = '@'
						fallingSpace[y].cells[x+1] = '.'
						fallingSpace[y+1].cells[x-1] = '@'
						fallingSpace[y+1].cells[x+1] = '.'
						x--
					}
				}
			}
		case right:
			if x+rock.getWidth() < 7 {
				switch rock {
				case stripeH:
					if fallingSpace[y].cells[x+4] == '.' {
						fallingSpace[y].cells[x+4] = '@'
						fallingSpace[y].cells[x] = '.'
						x++
					}
				case plus:
					if fallingSpace[y].cells[x+2] == '.' &&
						fallingSpace[y+1].cells[x+3] == '.' &&
						fallingSpace[y+2].cells[x+2] == '.' {
						fallingSpace[y].cells[x+2] = '@'
						fallingSpace[y].cells[x+1] = '.'
						fallingSpace[y+1].cells[x+3] = '@'
						fallingSpace[y+1].cells[x] = '.'
						fallingSpace[y+2].cells[x+2] = '@'
						fallingSpace[y+2].cells[x+1] = '.'
						x++
					}
				case lreverted:
					if fallingSpace[y].cells[x+3] == '.' &&
						fallingSpace[y+1].cells[x+3] == '.' &&
						fallingSpace[y+2].cells[x+3] == '.' {
						fallingSpace[y].cells[x+3] = '@'
						fallingSpace[y].cells[x] = '.'
						fallingSpace[y+1].cells[x+3] = '@'
						fallingSpace[y+1].cells[x+2] = '.'
						fallingSpace[y+2].cells[x+3] = '@'
						fallingSpace[y+2].cells[x+2] = '.'
						x++
					}
				case stripeV:
					if fallingSpace[y].cells[x+1] == '.' &&
						fallingSpace[y+1].cells[x+1] == '.' &&
						fallingSpace[y+2].cells[x+1] == '.' &&
						fallingSpace[y+3].cells[x+1] == '.' {
						fallingSpace[y].cells[x+1] = '@'
						fallingSpace[y].cells[x] = '.'
						fallingSpace[y+1].cells[x+1] = '@'
						fallingSpace[y+1].cells[x] = '.'
						fallingSpace[y+2].cells[x+1] = '@'
						fallingSpace[y+2].cells[x] = '.'
						fallingSpace[y+3].cells[x+1] = '@'
						fallingSpace[y+3].cells[x] = '.'
						x++
					}
				case cube:
					if fallingSpace[y].cells[x+2] == '.' &&
						fallingSpace[y+1].cells[x+2] == '.' {
						fallingSpace[y].cells[x+2] = '@'
						fallingSpace[y].cells[x] = '.'
						fallingSpace[y+1].cells[x+2] = '@'
						fallingSpace[y+1].cells[x] = '.'
						x++
					}
				}
			}
		}

		var rest = false
		if y == 0 {
			rest = true
		} else {
			switch rock {
			case stripeH:
				if fallingSpace[y-1].cells[x] == '.' &&
					fallingSpace[y-1].cells[x+1] == '.' &&
					fallingSpace[y-1].cells[x+2] == '.' &&
					fallingSpace[y-1].cells[x+3] == '.' {
					fallingSpace[y-1].cells[x] = '@'
					fallingSpace[y].cells[x] = '.'
					fallingSpace[y-1].cells[x+1] = '@'
					fallingSpace[y].cells[x+1] = '.'
					fallingSpace[y-1].cells[x+2] = '@'
					fallingSpace[y].cells[x+2] = '.'
					fallingSpace[y-1].cells[x+3] = '@'
					fallingSpace[y].cells[x+3] = '.'
					y--
				} else {
					rest = true
				}
			case plus:
				if fallingSpace[y].cells[x] == '.' &&
					fallingSpace[y-1].cells[x+1] == '.' &&
					fallingSpace[y].cells[x+2] == '.' {
					fallingSpace[y].cells[x] = '@'
					fallingSpace[y+1].cells[x] = '.'
					fallingSpace[y-1].cells[x+1] = '@'
					fallingSpace[y+2].cells[x+1] = '.'
					fallingSpace[y].cells[x+2] = '@'
					fallingSpace[y+1].cells[x+2] = '.'
					y--
				} else {
					rest = true
				}
			case lreverted:
				if fallingSpace[y-1].cells[x] == '.' &&
					fallingSpace[y-1].cells[x+1] == '.' &&
					fallingSpace[y-1].cells[x+2] == '.' {
					fallingSpace[y-1].cells[x] = '@'
					fallingSpace[y].cells[x] = '.'
					fallingSpace[y-1].cells[x+1] = '@'
					fallingSpace[y].cells[x+1] = '.'
					fallingSpace[y-1].cells[x+2] = '@'
					fallingSpace[y+2].cells[x+2] = '.'
					y--
				} else {
					rest = true
				}
			case stripeV:
				if fallingSpace[y-1].cells[x] == '.' {
					fallingSpace[y-1].cells[x] = '@'
					fallingSpace[y+3].cells[x] = '.'
					y--
				} else {
					rest = true
				}
			case cube:
				if fallingSpace[y-1].cells[x] == '.' &&
					fallingSpace[y-1].cells[x+1] == '.' {
					fallingSpace[y-1].cells[x] = '@'
					fallingSpace[y+1].cells[x] = '.'
					fallingSpace[y-1].cells[x+1] = '@'
					fallingSpace[y+1].cells[x+1] = '.'
					y--
				} else {
					rest = true
				}
			}
		}
		if rest {
			switch rock {
			case stripeH:
				fallingSpace[y].cells[x] = '#'
				fallingSpace[y].cells[x+1] = '#'
				fallingSpace[y].cells[x+2] = '#'
				fallingSpace[y].cells[x+3] = '#'
			case plus:
				fallingSpace[y].cells[x+1] = '#'
				fallingSpace[y+1].cells[x] = '#'
				fallingSpace[y+1].cells[x+1] = '#'
				fallingSpace[y+1].cells[x+2] = '#'
				fallingSpace[y+2].cells[x+1] = '#'
			case lreverted:
				fallingSpace[y].cells[x] = '#'
				fallingSpace[y].cells[x+1] = '#'
				fallingSpace[y].cells[x+2] = '#'
				fallingSpace[y+1].cells[x+2] = '#'
				fallingSpace[y+2].cells[x+2] = '#'
			case stripeV:
				fallingSpace[y].cells[x] = '#'
				fallingSpace[y+1].cells[x] = '#'
				fallingSpace[y+2].cells[x] = '#'
				fallingSpace[y+3].cells[x] = '#'
			case cube:
				fallingSpace[y].cells[x] = '#'
				fallingSpace[y].cells[x+1] = '#'
				fallingSpace[y+1].cells[x] = '#'
				fallingSpace[y+1].cells[x+1] = '#'
			}

		LOOP:
			for i := len(fallingSpace) - 1; i >= 0; i-- {
				for _, c := range fallingSpace[i].cells {
					if c == '#' {
						height = i + 1
						break LOOP
					}
				}
			}

			return
		}
	}
}

func main() {
	bs, _ := os.ReadFile("day17.txt")
	wind = strings.TrimSpace(string(bs))

	// part 1
	//for i := 0; i < 2022; i++ {
	//	simulate(false)
	//}
	//fmt.Println(height)

	// part 2
	timeStart := time.Now()
	var h = 0
	for i := 0; i < 1000000000000; i++ {
		if height > 10000 && height%2 == 0 {
			height /= 2
			h += height
			fallingSpace = fallingSpace[height:]
			fmt.Printf("%s: %.4f%% %d\n", time.Since(timeStart), float64(i)/1000000000000.0*100, h)
		}
		simulate(false)
	}
	h += height
	fmt.Println(h)
}
