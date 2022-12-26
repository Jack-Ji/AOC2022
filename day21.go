package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ops int

const (
	none Ops = iota
	add
	sub
	mul
	div
)

func (op Ops) apply(v1, v2 int) int {
	switch op {
	case add:
		return v1 + v2
	case sub:
		return v1 - v2
	case mul:
		return v1 * v2
	case div:
		return v1 / v2
	default:
		panic("unreachable")
	}
}

func strToOp(s string) Ops {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	}
	panic("unreachable")
}

type Monkey struct {
	name   string
	op     Ops
	v1, v2 string
	value  int
	final  bool
}

var mks = map[string]Monkey{}

func getRoot() int {
	for !mks["root"].final {
		for k, v := range mks {
			if v.final {
				continue
			}
			if !mks[v.v1].final || !mks[v.v2].final {
				continue
			}
			v.value = v.op.apply(mks[v.v1].value, mks[v.v2].value)
			v.final = true
			mks[k] = v
		}
	}

	return mks["root"].value
}

func getHumn() int {
	root := mks["root"]
	for {
		var merged int
		for k, v := range mks {
			if v.final {
				continue
			}
			if !mks[v.v1].final || !mks[v.v2].final {
				continue
			}
			v.value = v.op.apply(mks[v.v1].value, mks[v.v2].value)
			v.final = true
			mks[k] = v
			merged++
		}
		if merged == 0 {
			break
		}
	}

	if mks[root.v1].final {
		m2 := mks[root.v2]
		m2.final = true
		m2.value = mks[root.v1].value
		mks[root.v2] = m2
	} else {
		m1 := mks[root.v1]
		m1.final = true
		m1.value = mks[root.v2].value
		mks[root.v1] = m1
	}

	for !mks["humn"].final {
		for _, v := range mks {
			if v.final && v.op != none {
				if mks[v.v1].final && !mks[v.v2].final {
					m1 := mks[v.v1]
					m2 := mks[v.v2]
					switch v.op {
					case add:
						m2.value = v.value - m1.value
					case sub:
						m2.value = m1.value - v.value
					case mul:
						m2.value = v.value / m1.value
					case div:
						m2.value = m1.value / v.value
					default:
						panic("unreachable")
					}
					m2.final = true
					mks[v.v2] = m2
				} else if !mks[v.v1].final && mks[v.v2].final {
					m1 := mks[v.v1]
					m2 := mks[v.v2]
					switch v.op {
					case add:
						m1.value = v.value - m2.value
					case sub:
						m1.value = m2.value + v.value
					case mul:
						m1.value = v.value / m2.value
					case div:
						m1.value = m2.value * v.value
					default:
						panic("unreachable")
					}
					m1.final = true
					mks[v.v1] = m1
				}
			}
		}
	}
	return mks["humn"].value
}

func main() {
	bs, _ := os.ReadFile("day21.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	for _, s := range lines {
		var (
			md     = strings.TrimSpace(s)
			monkey Monkey
		)
		if len(md) > 10 {
			monkey = Monkey{
				name: md[:4],
				op:   strToOp(md[11:12]),
				v1:   md[6:10],
				v2:   md[13:],
			}
		} else {
			if md[:4] == "humn" {
				monkey = Monkey{
					name:  md[:4],
					final: false,
				}
			} else {
				value, _ := strconv.Atoi(md[6:])
				monkey = Monkey{
					name:  md[:4],
					value: value,
					final: true,
				}
			}
		}
		mks[monkey.name] = monkey
	}

	//part 1
	//fmt.Println(getRoot())

	//part 2
	fmt.Println(getHumn())
}
