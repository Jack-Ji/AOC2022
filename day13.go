package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Value struct {
	v *int // null means Value is a list
	l []*Value
}

func (value Value) String() string {
	var result string
	if value.v != nil {
		result += fmt.Sprintf("%d", *value.v)
	} else {
		result += "["
		for i, v := range value.l {
			if i > 0 {
				result += ","
			}
			result += v.String()
		}
		result += "]"
	}
	return result
}

func compare(v0, v1 *Value) (result int) {
	// both number
	if v0.v != nil && v1.v != nil {
		return *v0.v - *v1.v
	}

	// both list
	if v0.v == nil && v1.v == nil {
		var idx = 0
		for {
			if idx == len(v0.l) && idx == len(v1.l) {
				return 0
			}
			if idx < len(v0.l) && idx == len(v1.l) {
				return 1
			}
			if idx == len(v0.l) && idx < len(v1.l) {
				return -1
			}
			result := compare(v0.l[idx], v1.l[idx])
			if result == 0 {
				idx += 1
				continue
			}
			return result
		}
	}

	// one of them is list
	if v1.v == nil {
		if v1.l == nil {
			return 1
		}

		wrapped := Value{l: []*Value{v0}}
		return compare(&wrapped, v1)
	} else {
		if v0.l == nil {
			return -1
		}

		wrapped := Value{l: []*Value{v1}}
		return compare(v0, &wrapped)
	}
}

var vs []*Value

func getNext(s string) string {
	if s == "" || s == "]" {
		return ""
	}

	var end = 1
	if s[0] == '[' {
		var brackets = 1
		for brackets > 0 {
			switch s[end] {
			case '[':
				brackets += 1
			case ']':
				brackets -= 1
			}
			end += 1
		}
	} else {
		for {
			if !(s[end] >= '0' && s[end] <= '9') {
				break
			}
			end += 1
		}
	}
	return s[:end]
}

func parse(s string) *Value {
	if s == "" {
		panic("unreachable")
	}

	if s == "[]" {
		return &Value{}
	}

	singleValue, e := strconv.Atoi(s)
	if e == nil {
		return &Value{v: &singleValue}
	}

	multipleValues := &Value{}
	if s[0] != '[' {
		panic("unreachable")
	}
	var off = 1
	for {
		ss := getNext(s[off:])
		if len(ss) == 0 {
			break
		}
		multipleValues.l = append(multipleValues.l, parse(ss))
		off += len(ss) + 1
	}
	return multipleValues
}

func main() {
	bs, _ := os.ReadFile("day13.txt")
	pairs := strings.Split(strings.TrimSpace(string(bs)), "\r\n\r\n")

	for _, p := range pairs {
		ps := strings.Split(p, "\n")
		p1, p2 := strings.TrimSpace(ps[0]), strings.TrimSpace(ps[1])
		vs = append(vs, parse(p1), parse(p2))
	}

	var sum = 0
	for i := 0; i < len(vs); i += 2 {
		result := compare(vs[i], vs[i+1])
		if result == 0 {
			panic("unreachable")
		}
		if result < 0 {
			index := (i / 2) + 1
			sum += index
		}
	}
	fmt.Println(sum)

	// Sort packets
	sort.Slice(vs, func(i, j int) bool {
		return compare(vs[i], vs[j]) < 0
	})
	var divider0 = parse("[[2]]")
	var divider1 = parse("[[6]]")
	idx0, _ := sort.Find(len(vs), func(i int) int {
		return compare(divider0, vs[i])
	})
	idx1, _ := sort.Find(len(vs), func(i int) int {
		return compare(divider1, vs[i])
	})
	fmt.Println((idx0 + 1) * (idx1 + 2))
}
