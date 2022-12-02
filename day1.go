package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, _ := os.ReadFile("day1.txt")
	fs := strings.Split(string(bs), "\n")

	var (
		max1 = 0
		max2 = 0
		max3 = 0
		temp = 0
	)
	for i, f := range fs {
		sf := strings.TrimSpace(f)
		x, _ := strconv.Atoi(sf)
		temp += x

		if sf == "" || i == len(fs)-1 {
			if temp > max1 {
				max3 = max2
				max2 = max1
				max1 = temp
			} else if temp > max2 {
				max3 = max2
				max2 = temp
			} else if temp > max3 {
				max3 = temp
			}
			temp = 0
		}
	}
	fmt.Println(max1, max2, max3)
	fmt.Println(max1 + max2 + max3)
}
