package main

import (
	"fmt"
	"os"
	"strings"
)

type Valve struct {
	name    string
	tunnels []int
	rate    int

	shortests map[int]int // shortest distance to other valves
}

var (
	valves    []Valve
	rates     []int
	start     int
	threshold int
)

// Use BFS algorithm to calculate shortest distance between nodes
func calcShortests() {
	for i := range valves {
		var visited = map[int]bool{i: true}
		for len(visited) != len(valves) {
			var recentVisited = map[int]bool{}
			for k := range visited {
				for _, t := range valves[k].tunnels {
					if visited[t] {
						continue
					}
					recentVisited[t] = true
					distance := valves[i].shortests[k] + 1
					if valves[i].shortests[t] == 0 ||
						valves[i].shortests[t] > distance {
						valves[i].shortests[t] = distance
					}
				}
			}
			for k := range recentVisited {
				visited[k] = true
			}
		}
	}
}

// valve -> time left
type Solution map[int]int

func (sl Solution) clone() Solution {
	var cloned = make(Solution)
	for k, v := range sl {
		cloned[k] = v
	}
	return cloned
}

func (sl Solution) value() int {
	var value int
	for k, v := range sl {
		value += valves[k].rate * v
	}
	return value
}

// Get all possible solutions, with one man
func getSolutions1(idx int, halfSolution Solution, time int) (ss []Solution) {
	if time < 3 || len(halfSolution) == len(rates) {
		return []Solution{halfSolution}
	}

	for _, i := range rates {
		if halfSolution[i] == 0 {
			sl := halfSolution.clone()
			sl[i] = time - valves[idx].shortests[i] - 1
			ss = append(ss, getSolutions1(i, sl, sl[i])...)
		}
	}
	return ss
}

// Get all possible solutions, with two man
func getSolutions2(idx1, idx2 int, halfSolution Solution, time1, time2 int) (ss []Solution) {
	if (time1 < 3 && time2 < 3) || len(halfSolution) == len(rates) {
		if halfSolution.value() < threshold {
			return
		}
		return []Solution{halfSolution.clone()}
	}

	if time1 < 3 {
		for _, i := range rates {
			if halfSolution[i] == 0 {
				sl := halfSolution.clone()
				sl[i] = time2 - valves[idx2].shortests[i] - 1
				ss = append(ss, getSolutions2(idx1, i, sl, time1, sl[i])...)
			}
		}
	} else if time2 < 3 {
		for _, i := range rates {
			if halfSolution[i] == 0 {
				sl := halfSolution.clone()
				sl[i] = time1 - valves[idx1].shortests[i] - 1
				ss = append(ss, getSolutions2(i, idx2, sl, sl[i], time2)...)
			}
		}
	} else {
		var xs []int
		for _, i := range rates {
			if halfSolution[i] == 0 {
				xs = append(xs, i)
			}
		}
		if len(xs) == 1 {
			dst := xs[0]

			if time1-valves[idx1].shortests[dst]-1 >= 0 {
				sl := halfSolution.clone()
				halfSolution[dst] = time1 - valves[idx1].shortests[dst] - 1
				if sl.value() >= threshold {
					ss = append(ss, sl)
				}
			}

			if time2-valves[idx2].shortests[dst]-1 >= 0 {
				sl := halfSolution.clone()
				sl[dst] = time2 - valves[idx2].shortests[dst] - 1
				if sl.value() >= threshold {
					ss = append(ss, sl)
				}
			}
		} else {
			for i := 0; i < len(xs); i++ {
				dst1 := xs[i]
				if time1-valves[idx1].shortests[dst1]-1 < 0 {
					continue
				}
				halfSolution[dst1] = time1 - valves[idx1].shortests[dst1] - 1

				for j := 0; j < len(xs); j++ {
					if i == j {
						break
					}
					dst2 := xs[j]
					if time2-valves[idx2].shortests[dst2]-1 >= 0 {
						sl := halfSolution.clone()
						sl[dst2] = time2 - valves[idx2].shortests[dst2] - 1
						ss = append(ss, getSolutions2(dst1, dst2, sl, sl[dst1], sl[dst2])...)
					}
				}

				delete(halfSolution, dst1)
			}
		}
	}

	return ss
}

func main() {
	bs, _ := os.ReadFile("day16.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	for _, l := range lines {
		v := Valve{
			shortests: map[int]int{},
		}
		fmt.Sscanf(
			strings.TrimSpace(l),
			"Valve %s has flow rate=%d; tunnels lead to valves",
			&v.name, &v.rate,
		)
		valves = append(valves, v)
		if v.rate > 0 {
			rates = append(rates, len(valves)-1)
		}
		if v.name == "AA" {
			start = len(valves) - 1
		}
	}
	for i, l := range lines {
		line := strings.TrimSpace(l)
		idx := strings.Index(line, " to valve")
		ns := strings.Split(strings.TrimSpace(line[idx+len(" to valve")+1:]), ", ")
		for _, n := range ns {
			for j, v := range valves {
				if v.name == n {
					valves[i].tunnels = append(valves[i].tunnels, j)
					break
				}
			}
		}
	}

	calcShortests()

	// part1
	sls := getSolutions1(start, Solution{}, 30)
	fmt.Println("Part1, got", len(sls), "solutions")
	var (
		max  = 0
		sidx = 0
	)
	for i, sl := range sls {
		v := sl.value()
		if v > max {
			max = v
			sidx = i
		}
	}
	fmt.Println(sls[sidx], max)

	// part2
	threshold = max
	sls = getSolutions2(start, start, Solution{}, 26, 26)
	fmt.Println("Part2, got", len(sls), "solutions")
	max = 0
	sidx = 0
	for i, sl := range sls {
		v := sl.value()
		if v > max {
			max = v
			sidx = i
		}
	}
	fmt.Println(sls[sidx], max)
}
