package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const total_space = 70000000
const need_space = 30000000

var dirs = map[string]int{}

func updateSize(path string, size int) {
	if !strings.HasPrefix(path, "/") {
		panic("unreachable")
	}

	dirs["/"] = dirs["/"] + size

	if len(path) > 1 {
		dirs[path] = dirs[path] + size
		off := 1
		for {
			idx := strings.IndexByte(path[off:], '/')
			if idx < 0 {
				break
			}
			off += idx
			subPath := path[:off]
			dirs[subPath] = dirs[subPath] + size
			off += 1
		}
	}
}

func main() {
	bs, _ := os.ReadFile("day7.txt")
	fs := strings.Split(string(bs), "\n")

	var path string
	for _, f := range fs {
		line := strings.TrimSpace(f)
		if len(line) == 0 {
			break
		}
		if line[0] == '$' {
			if line[:4] == "$ cd" {
				switch cdpath := line[5:]; cdpath {
				case "/":
					path = "/"
				case "..":
					c := strings.LastIndexByte(path, '/')
					if c == 0 {
						path = "/"
					} else {
						path = path[:c]
					}
				default:
					if path == "/" {
						path = path + cdpath
					} else {
						path = path + "/" + cdpath
					}
				}
			}
			continue
		}

		if strings.HasPrefix(line, "dir ") {
			continue
		}
		size, _ := strconv.Atoi(line[:strings.IndexByte(line, ' ')])
		updateSize(path, size)
	}

	var (
		sum               = 0
		minimum_free_size = need_space - (total_space - dirs["/"])
		selected_size     = total_space
	)
	for _, s := range dirs {
		if s > minimum_free_size && s < selected_size {
			selected_size = s
		}
		if s <= 100000 {
			sum += s
		}
	}
	fmt.Println(sum, selected_size)
}
