package main

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/aocutils"
)

var min int = 0
var max int = 100

var answer int = 0
var current int = 50

func parseCommand(line string) int {
	if strings.HasPrefix(line, "L") {
		num, _ := strings.CutPrefix(line, "L")
		if s, err := strconv.Atoi(num); err == nil {
			aocutils.VPrintf("L : %d\n", -s)
			return -s
		}
	}
	if strings.HasPrefix(line, "R") {
		num, _ := strings.CutPrefix(line, "R")
		if s, err := strconv.Atoi(num); err == nil {
			aocutils.VPrintf("R : %d\n", s)
			return s
		}
	}
	return -1
}

func mod(a, n int) int {
	return (a%n + n) % n
}

// floorDiv100 returns floor(a / 100). Go's native integer division
// truncates toward zero, which gives the wrong answer for negative
// operands, so this can't just be "a / 100".
func floorDiv100(a int) int {
	if a >= 0 || a%100 == 0 {
		return a / 100
	}
	return a/100 - 1
}

// countZeroHits returns how many times the dial points at 0 while
// rotating by delta d starting from position p (0 <= p < 100). It
// counts every intermediate pass through 0 as well as the final
// landing, and works for any magnitude or direction of d.
func countZeroHits(p, d int) int {
	switch {
	case d > 0:
		return floorDiv100(p+d) - floorDiv100(p)
	case d < 0:
		return floorDiv100(p-1) - floorDiv100(p+d-1)
	default:
		return 0
	}
}

func main() {
	aocutils.Initiate()

	lineArr := aocutils.ReadLinesFromInputText("input.txt")

	for i, line := range lineArr {
		aocutils.VPrintf("=====Iteration Number : %d=====\n", i+1)
		aocutils.VPrintf("Start : %d\n", current)

		value := parseCommand(line)
		hits := countZeroHits(current, value)
		aocutils.VPrintf("Hits : %d\n", hits)

		current = mod(current+value-min, max-min) + min
		aocutils.VPrintf("Current : %d\n", current)

		answer += hits

		aocutils.VPrintf("Answer : %d\n", answer)
	}

	fmt.Printf("Answer is : %d", answer)
}
