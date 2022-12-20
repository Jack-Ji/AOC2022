package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Blueprint struct {
	id                int
	oreRobotCost      int    // consume ore
	clayRobotCost     int    // consume ore
	obsidianRobotCost [2]int // consume ore and clay
	geodeRobotCost    [2]int // consume ore and obsidian
}

func (b Blueprint) produceGeode(
	gas int, // gas left
	oreN, clayN, obsidianN int, // materials we have
	oreRobots, clayRobots, obsidianRobots, geodeRobots int, // robots we have
) (produced int) {
	if gas > 19 {
		fmt.Println(gas, "\tmaterials:", oreN, clayN, obsidianN,
			"\trobots:", oreRobots, clayRobots, obsidianRobots, geodeRobots)

		startTime := time.Now()
		defer func() {
			fmt.Println(gas, "\tconsumed:", time.Since(startTime), "\tproduced:", produced)
		}()
	}

	// no gas left
	if gas == 0 {
		produced = 0
		return
	}

	// no geode robot at all
	if gas == 1 && geodeRobots == 0 {
		produced = 0
		return
	}

	// no enough material to build at least one geode robot
	if gas == 2 && geodeRobots == 0 && (oreN < b.geodeRobotCost[0] || obsidianN < b.geodeRobotCost[1]) {
		produced = 0
		return
	}

	var (
		max              int
		tryNothing       bool
		tryOreRobot      bool
		tryClayRobot     bool
		tryObsidianRobot bool
		tryGeodeRobot    bool
	)
	for {
		var (
			ore               = oreN
			clay              = clayN
			obsidian          = obsidianN
			geode             int
			newOreRobots      int
			newClayRobots     int
			newObsidianRobots int
			newGeodeRobots    int
		)

		// decide what to build
		if !tryNothing {
			tryNothing = true
		} else if !tryOreRobot && ore >= b.oreRobotCost {
			tryOreRobot = true
			ore -= b.oreRobotCost
			newOreRobots++
		} else if !tryClayRobot && ore >= b.clayRobotCost {
			tryClayRobot = true
			ore -= b.clayRobotCost
			newClayRobots++
		} else if !tryObsidianRobot && ore >= b.obsidianRobotCost[0] && clay >= b.obsidianRobotCost[1] {
			tryObsidianRobot = true
			ore -= b.obsidianRobotCost[0]
			clay -= b.obsidianRobotCost[1]
			newObsidianRobots++
		} else if !tryGeodeRobot && ore >= b.geodeRobotCost[0] && obsidian >= b.geodeRobotCost[1] {
			tryGeodeRobot = true
			ore -= b.geodeRobotCost[0]
			obsidian -= b.geodeRobotCost[1]
			newGeodeRobots++
		} else {
			break
		}

		// collect robots' outputs
		ore += oreRobots
		clay += clayRobots
		obsidian += obsidianRobots
		geode = geodeRobots

		// progress further
		geode += b.produceGeode(gas-1, ore, clay, obsidian, oreRobots+newOreRobots,
			clayRobots+newClayRobots, obsidianRobots+newObsidianRobots, geodeRobots+newGeodeRobots)

		if geode > max {
			max = geode
		}
	}

	produced = max
	return max
}

var blueprints []Blueprint

func main() {
	bs, _ := os.ReadFile("day19.txt")
	lines := strings.Split(strings.TrimSpace(string(bs)), "\n")

	for _, s := range lines {
		b := Blueprint{}
		fmt.Sscanf(strings.TrimSpace(s), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.id, &b.oreRobotCost, &b.clayRobotCost, &b.obsidianRobotCost[0], &b.obsidianRobotCost[1], &b.geodeRobotCost[0], &b.geodeRobotCost[1])
		blueprints = append(blueprints, b)
	}

	// part1
	//var qualityLevels = make([]int, len(blueprints))
	//for i, b := range blueprints {
	//	fmt.Println("calculating #", b.id)
	//	qualityLevels[i] = b.produceGeode(24, 0, 0, 0, 1, 0, 0, 0) * b.id
	//}
	//var sum = 0
	//for _, lvl := range qualityLevels {
	//	sum += lvl
	//}
	//fmt.Println(sum)

	// part2
	var ns = [3]int{}
	for i, b := range blueprints[:3] {
		fmt.Println("calculating #", b.id)
		ns[i] = b.produceGeode(32, 0, 0, 0, 1, 0, 0, 0)
	}
	fmt.Println(ns[0] * ns[1] * ns[2])
}
