package main

import (
	"AdventOfCode/aocutils"
	"fmt"
)

func maxJoltage(bank string) int64 {
	n := len(bank)
	k := 12

	result := make([]byte, 0, k)
	start := 0

	for remaining := k; remaining > 0; remaining-- {
		// We must leave enough characters after the chosen one.
		end := n - remaining

		best := start

		for i := start; i <= end; i++ {
			if bank[i] > bank[best] {
				best = i
			}
		}

		result = append(result, bank[best])
		start = best + 1
	}

	var joltage int64
	for _, digit := range result {
		joltage = joltage*10 + int64(digit-'0')
	}

	return joltage
}

func main() {
	aocutils.Initiate()
	lines := aocutils.ReadLinesFromInputText("input.txt")

	var total int64

	for _, line := range lines {
		joltage := maxJoltage(line)
		total += joltage
	}

	fmt.Println("Total:", total)
}
