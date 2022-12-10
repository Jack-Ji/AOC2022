package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction int

const (
	noop Instruction = iota
	add
)

type Code struct {
	instruction Instruction
	operand     int
}

type CPU struct {
	cs  []Code
	pc  int
	reg int
}

func NewCPU(cs []Code) CPU {
	return CPU{
		cs:  cs,
		reg: 1,
	}
}

func (cpu *CPU) run(cycle int) int {
	var (
		c           = 1
		accumulateC = 0
	)
	for c < cycle && cpu.pc < len(cpu.cs) {
		code := cpu.cs[cpu.pc]
		switch code.instruction {
		case noop:
			cpu.pc++
		case add:
			accumulateC += 1
			if accumulateC == 2 {
				accumulateC = 0
				cpu.reg += code.operand
				cpu.pc++
			}
		}
		c += 1
	}
	return cpu.reg * cycle
}

func (cpu *CPU) runAll() {
	var (
		c           = 0
		accumulateC = 0
	)
	for cpu.pc < len(cpu.cs) {
		c += 1
		drawPos := c%40 - 1
		if drawPos < 0 {
			drawPos = 39
		}
		spritePos := cpu.reg
		if drawPos-spritePos >= -1 && drawPos-spritePos <= 1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if c%40 == 0 {
			fmt.Println()
		}

		code := cpu.cs[cpu.pc]
		switch code.instruction {
		case noop:
			cpu.pc++
		case add:
			accumulateC += 1
			if accumulateC == 2 {
				accumulateC = 0
				cpu.reg += code.operand
				cpu.pc++
			}
		}
	}
}

func (cpu *CPU) reset() {
	cpu.pc = 0
	cpu.reg = 1
}

func main() {
	bs, _ := os.ReadFile("day10.txt")
	fs := strings.Split(string(bs), "\n")

	var cs []Code
	for _, f := range fs {
		line := strings.TrimSpace(f)
		if len(line) == 0 {
			break
		}

		switch line[:4] {
		case "noop":
			cs = append(cs, Code{instruction: noop})
		case "addx":
			v, _ := strconv.Atoi(line[5:])
			cs = append(cs, Code{instruction: add, operand: v})
		}
	}

	cpu := NewCPU(cs)
	s1 := cpu.run(20)
	cpu.reset()
	s2 := cpu.run(60)
	cpu.reset()
	s3 := cpu.run(100)
	cpu.reset()
	s4 := cpu.run(140)
	cpu.reset()
	s5 := cpu.run(180)
	cpu.reset()
	s6 := cpu.run(220)
	fmt.Println(s1 + s2 + s3 + s4 + s5 + s6)

	cpu.reset()
	cpu.runAll()
}
