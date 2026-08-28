package aocutils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ReadLinesFromInputText reads every line from the given file
// (e.g. "input.txt" or "test.txt") and returns them as a slice of
// strings with line endings stripped.
func ReadLinesFromInputText(filename string) []string {
	fi, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open %s: %v", filename, err)
	}
	defer fi.Close()

	var lines []string
	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read %s: %v", filename, err)
	}

	return lines
}

func ReadLineFromSingleLineInputText(filename string, separator string) []string {
	lines := ReadLinesFromInputText(filename)
	if len(lines) != 1 {
		log.Fatalf("expected single line, got %d", len(lines))
	}
	var line []string = strings.Split(lines[0], separator)
	return line
}

func ReadLineFromSingleLineText(text string, separator string) []string {
	var line []string = strings.Split(text, separator)
	return line
}
