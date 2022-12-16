package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func abs(d int) int {
	if d >= 0 {
		return d
	}
	return -d
}

type Pair struct {
	sx, sy int
	bx, by int
}

func (p Pair) getCheckRange(y int) (int, int, bool) {
	distance := abs(p.sx-p.bx) + abs(p.sy-p.by)
	if p.sy-distance > y || p.sy+distance < y {
		return 0, 0, false
	}
	res := distance - abs(p.sy-y)
	return p.sx - res, p.sx + res, true
}

var pairs []Pair
var checkY = 2000000

func main() {
	bs, _ := os.ReadFile("day15.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	var inCheckLine = map[int]bool{}
	for _, l := range lines {
		p := Pair{}
		fmt.Sscanf(
			strings.TrimSpace(l),
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&p.sx, &p.sy, &p.bx, &p.by,
		)
		pairs = append(pairs, p)
		if p.by == checkY {
			inCheckLine[p.bx] = true
		}
	}

	/*
		// part 1
		var checkMap = map[int]bool{}
		for _, p := range pairs {
			rx0, rx1, result := p.getCheckRange(checkY)
			if !result {
				continue
			}
			for x := rx0; x <= rx1; x++ {
				if inCheckLine[x] {
					continue
				}
				checkMap[x] = true
			}
		}
		fmt.Println(len(checkMap))
	*/

	const limit = 4000000
	const jobN = 16
	var sum int64
	for i := 0; i <= limit; i++ {
		sum += int64(i + 1)
	}

	var (
		wg   sync.WaitGroup
		done = make(chan interface{})
	)
	for i := 0; i < jobN; i++ {
		wg.Add(1)

		go func(start, end int) {
			defer wg.Done()

		LOOP:
			for y := start; y <= end; y++ {
				var (
					xs    = [limit + 1]bool{false}
					total = sum
				)

				if y%1000 == 0 {
					select {
					case <-done:
						return
					default:
					}

					fmt.Printf("Searching line %d [%d-%d]\n", y, start, end)
				}
				for _, p := range pairs {
					rx0, rx1, result := p.getCheckRange(y)
					if !result {
						continue
					}
					if rx0 < 0 {
						rx0 = 0
					}
					if rx1 > limit {
						rx1 = limit
					}
					for x := rx0; x <= rx1; x++ {
						if !xs[x] {
							xs[x] = true
							total -= int64(x + 1)
							if total == 0 {
								continue LOOP
							}
						}
					}
				}
				if total > 0 {
					x := total - 1
					close(done)
					fmt.Println("FOUND it!", x, y, x*4000000+int64(y))
					return
				}
			}
		}(i*limit/jobN, (i+1)*limit/jobN)
	}

	wg.Wait()
}
