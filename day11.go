package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	items    []int
	ops      func(old int) int
	transfer func(v int) int
}

var monkeys = []Monkey{
	Monkey{
		items: []int{98, 89, 52},
		ops: func(old int) int {
			return old * 2
		},
		transfer: func(v int) int {
			if (v % 5) == 0 {
				return 6
			} else {
				return 1
			}
		},
	},
	Monkey{
		items: []int{57, 95, 80, 92, 57, 78},
		ops: func(old int) int {
			return old * 13
		},
		transfer: func(v int) int {
			if (v % 2) == 0 {
				return 2
			} else {
				return 6
			}
		},
	},
	Monkey{
		items: []int{82, 74, 97, 75, 51, 92, 83},
		ops: func(old int) int {
			return old + 5
		},
		transfer: func(v int) int {
			if (v % 19) == 0 {
				return 7
			} else {
				return 5
			}
		},
	},
	Monkey{
		items: []int{97, 88, 51, 68, 76},
		ops: func(old int) int {
			return old + 6
		},
		transfer: func(v int) int {
			if (v % 7) == 0 {
				return 0
			} else {
				return 4
			}
		},
	},
	Monkey{
		items: []int{63},
		ops: func(old int) int {
			return old + 1
		},
		transfer: func(v int) int {
			if (v % 17) == 0 {
				return 0
			} else {
				return 1
			}
		},
	},
	Monkey{
		items: []int{94, 91, 51, 63},
		ops: func(old int) int {
			return old + 4
		},
		transfer: func(v int) int {
			if (v % 13) == 0 {
				return 4
			} else {
				return 3
			}
		},
	},
	Monkey{
		items: []int{61, 54, 94, 71, 74, 68, 98, 83},
		ops: func(old int) int {
			return old + 2
		},
		transfer: func(v int) int {
			if (v % 3) == 0 {
				return 2
			} else {
				return 7
			}
		},
	},
	Monkey{
		items: []int{90, 56},
		ops: func(old int) int {
			return old * old
		},
		transfer: func(v int) int {
			if (v % 11) == 0 {
				return 3
			} else {
				return 5
			}
		},
	},
}

func main() {
	var inspectItems = [8]int{0}
	var mod = 5 * 2 * 19 * 7 * 17 * 13 * 3 * 11

	for i := 0; i < 10000; i++ {
		for j, m := range monkeys {
			for _, v := range m.items {
				nv := m.ops(v)
				nv %= mod
				dst := m.transfer(nv)
				if dst == j {
					panic("unreachable")
				}
				monkeys[dst].items = append(monkeys[dst].items, nv)
				inspectItems[j] += 1
			}
			monkeys[j].items = []int{}
		}
	}
	sort.Slice(inspectItems[:], func(i, j int) bool {
		return inspectItems[i] > inspectItems[j]
	})
	top1, top2 := inspectItems[0], inspectItems[1]
	fmt.Println(top1, top2, top1*top2)
}
