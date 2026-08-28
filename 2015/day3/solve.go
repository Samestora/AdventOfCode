package main

import (
	"AdventOfCode/aocutils"
	"fmt"
)

func parseCommand(line []string) (int, []int) {
	coords := make(map[[2]int]bool)
	coords[[2]int{0, 0}] = true

	current_pos := [2]int{0, 0}

	for _, char := range line {
		switch char {
		case "^": // y+
			current_pos[1]++
			coords[current_pos] = true
		case "v": // y-
			current_pos[1]--
			coords[current_pos] = true
		case ">": // x+
			current_pos[0]++
			coords[current_pos] = true
		case "<": // x-
			current_pos[0]--
			coords[current_pos] = true
		}
		aocutils.VPrintf("Command : %s", char)
		aocutils.VPrintf("Current Pos : %v", current_pos)
	}
	coord_list := make([]int, 0, len(coords))
	for coord := range coords {
		coord_list = append(coord_list, coord[0], coord[1])
	}
	return len(coords), coord_list
}

func parseRobotCommand(line []string) (int, []int) {
	coords := make(map[[2]int]bool)
	coords[[2]int{0, 0}] = true

	santa_pos := [2]int{0, 0}
	robot_pos := [2]int{0, 0}

	for idx, char := range line {
		if idx%2 == 0 {
			switch char {
			case "^": // y+
				santa_pos[1]++
			case "v": // y-
				santa_pos[1]--
			case ">": // x+
				santa_pos[0]++
			case "<": // x-
				santa_pos[0]--
			}
			coords[santa_pos] = true
		} else {
			switch char {
			case "^": // y+
				robot_pos[1]++
			case "v": // y-
				robot_pos[1]--
			case ">": // x+
				robot_pos[0]++
			case "<": // x-
				robot_pos[0]--
			}
			coords[robot_pos] = true
		}
	}
	coord_list := make([]int, 0, len(coords))
	for coord := range coords {
		coord_list = append(coord_list, coord[0], coord[1])
	}
	return len(coords), coord_list
}

func main() {
	aocutils.Initiate()
	line := aocutils.ReadLineFromSingleLineInputText("input.txt", "")
	answer1, coord_list := parseCommand(line)
	answer2, coord_list := parseRobotCommand(line)

	fmt.Printf("Part I : %d\n", answer1)
	fmt.Printf("Part II : %d\n", answer2)

	aocutils.VPrintf("Coord List : %v", coord_list)
}
