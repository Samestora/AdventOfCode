package main

import (
	"AdventOfCode/aocutils"
	"fmt"
	"strconv"
	"strings"
)

var total int

func chunking(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}

	for chunkSize := n / 2; chunkSize >= 1; chunkSize-- {
		// check if the chunk is evenly divisible
		if n%chunkSize != 0 {
			continue
		}

		// since it's perfect chunk duplication
		// like 123 is like 123123123 will counts as valid
		// but 123 if it's like 123123124 will counts as invalid EVEN if 123 duplicate at least twice
		pattern := s[0:chunkSize]
		isPerfect := false

		for i := chunkSize; i < n; i += chunkSize {
			if s[i:i+chunkSize] != pattern {
				isPerfect = false
				break
			}
			isPerfect = true
		}

		if isPerfect {
			return true
		}
	}
	return false
}

func parseCommand(command string) {
	ranges := aocutils.ReadLineFromSingleLineText(command, "-")

	if strings.HasPrefix(ranges[0], "0") || strings.HasPrefix(ranges[1], "0") {
		return
	}

	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		return
	}
	end, err := strconv.Atoi(ranges[1])
	if err != nil {
		return
	}

	for i := start; i <= end; i++ {
		aocutils.VPrintf("%d", i)
		str := strconv.Itoa(i)
		if chunking(str) {
			aocutils.VPrintf("Match: %d", i)
			total += i
		}
	}
}

func main() {
	aocutils.Initiate()

	lineArr := aocutils.ReadLineFromSingleLineInputText("input.txt", ",")

	for i, line := range lineArr {
		aocutils.VPrintf("======[%d: %s]======", i, line)
		parseCommand(line)
	}

	fmt.Printf("Total: %d", total)
}
