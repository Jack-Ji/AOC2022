package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seq []int
var indices []int

func getNthAfterPos(pos, nth int) int {
	return seq[(pos+nth)%len(seq)]
}

func getMixPosition(idx, step int) int {
	pos := (idx + step) % (len(seq) - 1)
	if pos < 0 {
		pos += len(seq) - 1
	}
	return pos
}

func mix(iidx int) {
	curPosition := indices[iidx]
	n := seq[curPosition]
	newPostion := getMixPosition(curPosition, n)
	if newPostion == curPosition {
		return
	}

	if newPostion > curPosition {
		copy(seq[curPosition:], seq[curPosition+1:newPostion+1])
		for i := iidx; i < len(indices); i++ {
			if indices[i] > curPosition && indices[i] <= newPostion {
				indices[i] -= 1
			}
		}
	} else {
		copy(seq[newPostion+1:], seq[newPostion:curPosition])
		for i := iidx; i < len(indices); i++ {
			if indices[i] >= newPostion && indices[i] < curPosition {
				indices[i] += 1
			}
		}
	}
	seq[newPostion] = n
	indices[iidx] = curPosition
}

func mix2(iidx int) {
	curPosition := indices[iidx]
	n := seq[curPosition]
	newPostion := getMixPosition(curPosition, n)
	if newPostion == curPosition {
		return
	}

	if newPostion > curPosition {
		copy(seq[curPosition:], seq[curPosition+1:newPostion+1])
		for i := 0; i < len(indices); i++ {
			if indices[i] > curPosition && indices[i] <= newPostion {
				indices[i] -= 1
			}
		}
	} else {
		copy(seq[newPostion+1:], seq[newPostion:curPosition])
		for i := 0; i < len(indices); i++ {
			if indices[i] >= newPostion && indices[i] < curPosition {
				indices[i] += 1
			}
		}
	}
	seq[newPostion] = n
	indices[iidx] = newPostion
}

func main() {
	bs, _ := os.ReadFile("day20.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	for i, s := range lines {
		n, _ := strconv.Atoi(strings.TrimSpace(s))
		seq = append(seq, n)
		indices = append(indices, i)
	}

	// part1
	//for i := range indices {
	//	mix(i)
	//}

	// part2
	for i := range seq {
		seq[i] = seq[i] * 811589153
	}
	for n := 10; n > 0; n-- {
		for i := range indices {
			mix2(i)
		}
	}

	var zidx int
	for i := range seq {
		if seq[i] == 0 {
			zidx = i
			break
		}
	}
	fmt.Println(
		getNthAfterPos(zidx, 1000),
		getNthAfterPos(zidx, 2000),
		getNthAfterPos(zidx, 3000),
		getNthAfterPos(zidx, 1000)+getNthAfterPos(zidx, 2000)+getNthAfterPos(zidx, 3000),
	)
}
