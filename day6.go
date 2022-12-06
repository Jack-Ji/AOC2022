package main

import (
	"fmt"
	"os"
)

func isDifferent(bs []byte) bool {
	mp := map[byte]bool{}
	for _, c := range bs {
		mp[c] = true
	}
	return len(mp) == len(bs)
}

func main() {
	bs, _ := os.ReadFile("day6.txt")

	for i := 3; i < len(bs); i++ {
		if isDifferent(bs[i-3 : i+1]) {
			fmt.Println(i + 1)
			break
		}
	}

	for i := 13; i < len(bs); i++ {
		if isDifferent(bs[i-13 : i+1]) {
			fmt.Println(i + 1)
			break
		}
	}
}
