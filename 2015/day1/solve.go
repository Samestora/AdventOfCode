package main

import "AdventOfCode/aocutils"

// Starts from 0
// ( +1
// ) -1

func parseCommand(char string) int {
	if char == "(" {
		return 1
	}
	return -1
}

func main() {
	aocutils.Initiate()

	var floor int

	line := aocutils.ReadLineFromSingleLineInputText("input.txt", "")

	for idx, char := range line {
		floor += parseCommand(char)
		if floor == -1 {
			println("Answer Pt II : ", idx+1)
			return
		}
	}

	println("Answer Pt I : ", floor)
}
