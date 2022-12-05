package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRanges(line string) ([]int, []int) {
	idx := strings.Index(line, ",")
	ridx1 := strings.Index(string(line[:idx]), "-")
	ridx2 := strings.Index(string(line[idx+1:]), "-")
	d1, _ := strconv.Atoi(line[:idx][:ridx1])
	d2, _ := strconv.Atoi(line[:idx][ridx1+1:])
	d3, _ := strconv.Atoi(line[idx+1:][:ridx2])
	d4, _ := strconv.Atoi(line[idx+1:][ridx2+1:])
	return []int{d1, d2}, []int{d3, d4}
}

func main() {
	bs, _ := os.ReadFile("day4.txt")
	lines := strings.Split(string(bs), "\n")

	var (
		count1 = 0
		count2 = 0
	)
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			break
		}
		r1, r2 := getRanges(l)

		if (r1[0] <= r2[0] && r1[1] >= r2[1]) ||
			(r2[0] <= r1[0] && r2[1] >= r1[1]) {
			count1 += 1
		}

		if (r1[0] >= r2[0] && r1[0] <= r2[1]) ||
			(r2[0] >= r1[0] && r2[0] <= r1[1]) {
			count2 += 1
		}

	}
	fmt.Println(count1, count2)
}
