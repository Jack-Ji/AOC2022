package main

import (
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	crates []byte
}

func (s *Stack) push(c byte) {
	s.crates = append(s.crates, c)
}

func (s *Stack) pushl(c byte) {
	cs := []byte{c}
	s.crates = append(cs, s.crates...)
}

func (s *Stack) pop() byte {
	if len(s.crates) == 0 {
		panic("unreachable")
	}
	popped := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return popped
}

func (s *Stack) pushn(cs []byte) {
	for _, c := range cs {
		s.push(c)
	}
}

func (s *Stack) pushn1(cs []byte) {
	for i := len(cs) - 1; i >= 0; i-- {
		s.push(cs[i])
	}
}

func (s *Stack) popn(n int) []byte {
	popped := []byte{}
	for n > 0 {
		popped = append(popped, s.pop())
		n -= 1
	}
	return popped
}

func (s *Stack) top() byte {
	return s.crates[len(s.crates)-1]
}

func main() {
	bs, _ := os.ReadFile("day5.txt")
	lines := strings.Split(string(bs), "\n")

	var stacks [9]Stack
	var idx int
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			idx = i
			break
		}

		for j := 0; j < 9; j++ {
			c := line[j*4+1]
			if c >= 'A' && c <= 'Z' {
				stacks[j].pushl(c)
			}
		}
	}

	var (
		cnt  int
		from int
		to   int
	)
	for _, line := range lines[idx+1:] {
		l := strings.TrimSpace(line)
		if l == "" {
			break
		}

		fmt.Sscanf(l, "move %d from %d to %d", &cnt, &from, &to)
		//stacks[to-1].pushn(stacks[from-1].popn(cnt))
		stacks[to-1].pushn1(stacks[from-1].popn(cnt))
	}
	fmt.Println(strings.Join([]string{
		string(stacks[0].top()),
		string(stacks[1].top()),
		string(stacks[2].top()),
		string(stacks[3].top()),
		string(stacks[4].top()),
		string(stacks[5].top()),
		string(stacks[6].top()),
		string(stacks[7].top()),
		string(stacks[8].top()),
	}, ""))
}
