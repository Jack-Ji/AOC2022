package main

import (
	"fmt"
	"os"
	"strings"
)

var sc = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
}

func score1(a, b byte) int {
	switch a {
	case 'A':
		switch b {
		case 'X':
			return 3 + sc['A']
		case 'Y':
			return 6 + sc['B']
		case 'Z':
			return 0 + sc['C']
		}
	case 'B':
		switch b {
		case 'X':
			return 0 + sc['A']
		case 'Y':
			return 3 + sc['B']
		case 'Z':
			return 6 + sc['C']
		}
	case 'C':
		switch b {
		case 'X':
			return 6 + sc['A']
		case 'Y':
			return 0 + sc['B']
		case 'Z':
			return 3 + sc['C']
		}
	}
	panic("unreachable")
}

func score2(a, b byte) int {
	switch b {
	case 'X':
		switch a {
		case 'A':
			return sc['C']
		case 'B':
			return sc['A']
		case 'C':
			return sc['B']
		}
	case 'Y':
		switch a {
		case 'A':
			return 3 + sc['A']
		case 'B':
			return 3 + sc['B']
		case 'C':
			return 3 + sc['C']
		}
	case 'Z':
		switch a {
		case 'A':
			return 6 + sc['B']
		case 'B':
			return 6 + sc['C']
		case 'C':
			return 6 + sc['A']
		}
	}
	panic("unreachable")
}

func main() {
	bs, _ := os.ReadFile("day2.txt")
	lines := strings.Split(string(bs), "\n")

	var (
		sum1 = 0
		sum2 = 0
	)
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			break
		}
		a := line[0]
		b := line[2]
		sum1 += score1(a, b)
		sum2 += score2(a, b)
	}
	fmt.Println(sum1, sum2)
}
