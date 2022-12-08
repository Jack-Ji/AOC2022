package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	grid   [][]byte
	width  int
	height int
)

type ByteSlice []byte

func (x ByteSlice) Len() int           { return len(x) }
func (x ByteSlice) Less(i, j int) bool { return x[i] > x[j] }
func (x ByteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func isVisible(x, y int) bool {
	if x == 0 || x == width-1 ||
		y == 0 || y == height-1 {
		return true
	}

	val := grid[y][x]

	// left part of row
	var bs []byte
	bs = append(bs, grid[y][:x+1]...)
	sort.Sort(ByteSlice(bs))
	if val == bs[0] && val > bs[1] {
		return true
	}

	// right part of row
	bs = bs[:0]
	bs = append(bs, grid[y][x:]...)
	sort.Sort(ByteSlice(bs))
	if val == bs[0] && val > bs[1] {
		return true
	}

	// up part of column
	bs = bs[:0]
	for i := 0; i <= y; i++ {
		bs = append(bs, grid[i][x])
	}
	sort.Sort(ByteSlice(bs))
	if val == bs[0] && val > bs[1] {
		return true
	}

	// down part of column
	bs = bs[:0]
	for i := y; i < height; i++ {
		bs = append(bs, grid[i][x])
	}
	sort.Sort(ByteSlice(bs))
	if val == bs[0] && val > bs[1] {
		return true
	}

	return false
}

func calcSceneScore(x, y int) int {
	if x == 0 || x == width-1 ||
		y == 0 || y == height-1 {
		return 0
	}

	val := grid[y][x]

	var left_score = 0
	for i := x - 1; i >= 0; i-- {
		left_score++
		if grid[y][i] >= val {
			break
		}
	}
	if left_score == 0 {
		return 0
	}

	var right_score = 0
	for i := x + 1; i < width; i++ {
		right_score++
		if grid[y][i] >= val {
			break
		}
	}
	if right_score == 0 {
		return 0
	}

	var up_score = 0
	for i := y - 1; i >= 0; i-- {
		up_score++
		if grid[i][x] >= val {
			break
		}
	}
	if up_score == 0 {
		return 0
	}

	var down_score = 0
	for i := y + 1; i < height; i++ {
		down_score++
		if grid[i][x] >= val {
			break
		}
	}
	if down_score == 0 {
		return 0
	}

	return left_score * right_score * up_score * down_score
}

func main() {
	bs, _ := os.ReadFile("day8.txt")
	fs := strings.Split(string(bs), "\n")

	for _, f := range fs {
		line := strings.TrimSpace(f)
		if len(line) == 0 {
			break
		}
		grid = append(grid, []byte(line))
	}

	width = len(grid[0])
	height = len(grid)

	var count int
	var biggest int
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if isVisible(i, j) {
				count++
			}

			score := calcSceneScore(i, j)
			if score > biggest {
				biggest = score
			}
		}
	}
	fmt.Println(count, biggest)
}
