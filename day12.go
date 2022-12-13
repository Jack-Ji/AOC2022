package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	none Direction = iota
	left
	right
	up
	down
)

func (d Direction) String() string {
	var s string
	switch d {
	case none:
		s = "*"
	case left:
		s = "<"
	case right:
		s = ">"
	case up:
		s = "^"
	case down:
		s = "v"
	}
	return s
}

type Position struct {
	h        byte
	x, y     int
	dir      Direction
	distance int
	selected bool
}

type Neighbor struct {
	idx int
	dir Direction
}

func (p Position) getIndex() int {
	return xyToIdx(p.x, p.y)
}

func (p Position) isNext(n *Position) bool {
	if p.h+1 == n.h || p.h >= n.h {
		return true
	}
	return false
}

func (p Position) getNeighbors() []Neighbor {
	var ns []Neighbor
	if p.x > 0 {
		n := &heightMap[p.y][p.x-1]
		if p.isNext(n) {
			ns = append(ns, Neighbor{n.getIndex(), right})
		}
	}
	if p.x < width-1 {
		n := &heightMap[p.y][p.x+1]
		if p.isNext(n) {
			ns = append(ns, Neighbor{n.getIndex(), left})
		}
	}
	if p.y > 0 {
		n := &heightMap[p.y-1][p.x]
		if p.isNext(n) {
			ns = append(ns, Neighbor{n.getIndex(), down})
		}
	}
	if p.y < height-1 {
		n := &heightMap[p.y+1][p.x]
		if p.isNext(n) {
			ns = append(ns, Neighbor{n.getIndex(), up})
		}
	}
	return ns
}

var (
	heightMap [][]Position
	width     int
	height    int
	xS, yS    int
	xE, yE    int
)

func xyToIdx(x, y int) int {
	return x + y*width
}

func idxToXY(idx int) (int, int) {
	return idx % width, idx / width
}

// BSF searching algorithm
func findPath(srcIdx, distanceThreshold int) {
	for i := range heightMap {
		for j := range heightMap[i] {
			heightMap[i][j].dir = none
			heightMap[i][j].distance = 0
		}
	}

	x, y := idxToXY(srcIdx)

	var visited = map[int]bool{xyToIdx(x, y): true}
	for {
		var oldVisitedNum = len(visited)
		var recentVisited = map[int]bool{}
		for k := range visited {
			vx, vy := idxToXY(k)
			ns := heightMap[vy][vx].getNeighbors()
			distance := heightMap[vy][vx].distance + 1
			if distance > distanceThreshold {
				// The path is too long, no need to search anymore
				return
			}
			for _, n := range ns {
				nx, ny := idxToXY(n.idx)
				if visited[n.idx] {
					continue
				}
				if recentVisited[n.idx] {
					if distance < heightMap[ny][nx].distance {
						heightMap[ny][nx].dir = n.dir
						heightMap[ny][nx].distance = distance
					} else {
						continue
					}
				} else {
					recentVisited[n.idx] = true
					heightMap[ny][nx].dir = n.dir
					heightMap[ny][nx].distance = distance
				}
			}
		}
		for k := range recentVisited {
			visited[k] = true
		}
		if len(visited) == oldVisitedNum {
			break
		}
	}
}

func main() {
	bs, _ := os.ReadFile("day12.txt")
	fs := strings.Split(strings.TrimSpace(string(bs)), "\n")
	height = len(fs)
	width = len(strings.TrimSpace(fs[0]))

	var candidates []int
	for j, f := range fs {
		line := strings.TrimSpace(f)

		var row []Position
		for i, b := range line {
			h := byte(b)
			if h == 'S' {
				h = 'a'
			}
			if h == 'E' {
				h = 'z'
			}
			row = append(row, Position{
				h:        h,
				x:        i,
				y:        j,
				dir:      none,
				distance: 0,
				selected: false,
			})
			if b == 'S' {
				xS, yS = i, j
			}
			if b == 'E' {
				xE, yE = i, j
			}
			if h == 'a' {
				candidates = append(candidates, xyToIdx(i, j))
			}
		}
		heightMap = append(heightMap, row)
	}

	var (
		shortestSpot = 0
		distance     = 1000
	)
	for i, idx := range candidates {
		findPath(idx, distance)
		if heightMap[yE][xE].distance == 0 {
			fmt.Println("Searching ", i+1, "/", len(candidates), ", distance: n/a")
			continue
		}
		fmt.Println("Searching ", i+1, "/", len(candidates), ", distance: ", heightMap[yE][xE].distance)
		if heightMap[yE][xE].distance < distance {
			shortestSpot = idx
			distance = heightMap[yE][xE].distance
		}
	}
	findPath(shortestSpot, distance)
	fmt.Println()

	var srcX, srcY = idxToXY(shortestSpot)
	var idxE = xyToIdx(xE, yE)
	var path = []int{idxE}
	for {
		x, y := idxToXY(path[len(path)-1])
		heightMap[y][x].selected = true
		if x == srcX && y == srcY {
			break
		}

		var next int
		switch heightMap[y][x].dir {
		case left:
			next = heightMap[y][x-1].getIndex()
		case right:
			next = heightMap[y][x+1].getIndex()
		case up:
			next = heightMap[y-1][x].getIndex()
		case down:
			next = heightMap[y+1][x].getIndex()
		default:
			panic("unreachable")
		}
		path = append(path, next)
	}

	fmt.Println(heightMap[yE][xE].distance)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			p := heightMap[i][j]
			if p.x == xS && p.y == yS {
				fmt.Printf("S")
			} else if p.x == xE && p.y == yE {
				fmt.Printf("E")
			} else if p.selected {
				fmt.Printf("%s", p.dir)
			} else {
				fmt.Printf("%c", p.h)
			}
		}
		fmt.Println()
	}
}
