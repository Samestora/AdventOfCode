package main

import (
	"AdventOfCode/aocutils"
	"strconv"
)

func wrappingPaper(l, w, h int) int {
	a := l
	b := w
	c := h

	if a > b {
		a, b = b, a
	} // Ensure a <= b
	if b > c {
		b, c = c, b
	} // Ensure b <= c
	if a > b {
		a, b = b, a
	} // Re-check a <= b after c moved

	return (2*l*w + 2*w*h + 2*h*l) + (a * b)
}

func ribbon(l, w, h int) int {
	a := l
	b := w
	c := h

	if a > b {
		a, b = b, a
	} // Ensure a <= b
	if b > c {
		b, c = c, b
	} // Ensure b <= c
	if a > b {
		a, b = b, a
	} // Re-check a <= b after c moved

	return (l * w * h) + (2*a + 2*b)
}

func main() {
	aocutils.Initiate()
	var answerWrap int = 0
	var answerRibbon int = 0
	lines := aocutils.ReadLinesFromInputText("input.txt")
	for _, line := range lines {
		dimension := aocutils.ReadLineFromSingleLineText(line, "x")
		l, _ := strconv.Atoi(dimension[0])
		w, _ := strconv.Atoi(dimension[1])
		h, _ := strconv.Atoi(dimension[2])

		answerWrap += wrappingPaper(l, w, h)
		answerRibbon += ribbon(l, w, h)
	}
	println("Answer Pt I : ", answerWrap)
	println("Answer Pt II : ", answerRibbon)
}
