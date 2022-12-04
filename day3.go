package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriority(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else {
		return int(c - 'A' + 27)
	}
}

func getSame(s string) byte {
	l := len(s) / 2
	for _, c1 := range s[:l] {
		for _, c2 := range s[l:] {
			if c1 == c2 {
				return byte(c1)
			}
		}
	}
	panic("unreachable")
}

func getSame3(s1, s2, s3 string) byte {
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				for _, c3 := range s3 {
					if c1 == c3 {
						return byte(c1)
					}
				}
			}
		}
	}
	panic("unreachable")
}

func main() {
	bs, _ := os.ReadFile("day3.txt")
	lines := strings.Split(string(bs), "\n")

	var (
		sum = 0
	)
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			break
		}
		sum += getPriority(getSame(l))
	}
	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(lines)/3; i += 1 {
		l1 := strings.TrimSpace(lines[i*3])
		l2 := strings.TrimSpace(lines[i*3+1])
		l3 := strings.TrimSpace(lines[i*3+2])
		sum += getPriority(getSame3(l1, l2, l3))
	}
	fmt.Println(sum)
}
