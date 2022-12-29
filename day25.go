package main

import (
	"fmt"
	"os"
	"strings"
)

func pow(a, b int) int {
	if b == 0 {
		return 1
	}

	var result = 1
	for b > 0 {
		b--
		result *= a
	}
	return result
}

func snafu(c byte) int {
	switch c {
	case '=':
		return -2
	case '-':
		return -1
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	}
	panic("unrechable")
}

func revert(s []byte) []byte {
	var i, j = 0, len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return s
}

func decimalToSnafu(n int) []byte {
	var (
		size   = 1
		result = make([]byte, 100)
	)

	result[0] = '0'
	for n > 0 {
		n--

		var off = 0
	LOOP:
		for {
			switch result[off] {
			case '=':
				result[off] = '-'
				break LOOP
			case '-':
				result[off] = '0'
				break LOOP
			case '0':
				result[off] = '1'
				break LOOP
			case '1':
				result[off] = '2'
				break LOOP
			case '2':
				result[off] = '='
				off++
				if off == size {
					size++
					result[off] = '1'
					break LOOP
				}
			default:
				panic("unrechable")
			}
		}
	}

	return revert(result[:size])
}

func snafuToDecimal(s []byte) int {
	var result int
	for i := 0; i < len(s); i++ {
		result += pow(5, len(s)-1-i) * snafu(s[i])
	}
	return result
}

func main() {
	bs, _ := os.ReadFile("day25.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	var sum int
	for _, s := range lines {
		n := strings.TrimSpace(s)
		sum += snafuToDecimal([]byte(n))
	}

	// part1
	fmt.Println("sum:", sum)
	fmt.Println("SNAFU:", string(decimalToSnafu(sum)))
}
