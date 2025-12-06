package main

import (
	"fmt"
	"iter"
	"os"
	"path"
	"strings"
	"time"
)

const (
	inputsDir = "inputs"
)

func main() {
	puzzles := []puzzle{
		&Day1{},
		&Day2{},
		&Day3{},
		&Day4{},
		&Day5{},
		&Day6{},
	}

	start := time.Now()

	for i, puzzle := range puzzles {
		day := i + 1
		inputPath := path.Join(inputsDir, fmt.Sprintf("day%d.txt", day))
		inputBytes, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("failed to read %s: %s\n", inputPath, err)
			continue
		}

		puzzle.Parse(inputBytes)

		fmt.Printf("Day %d Part 1: %d\n", day, puzzle.Part1())
		fmt.Printf("Day %d Part 2: %d\n", day, puzzle.Part2())
	}

	d := time.Since(start)

	fmt.Printf("Did %d puzzles in %s\n", len(puzzles), d)
}

type puzzle interface {
	Parse([]byte)
	Part1() int
	Part2() int
}

func Lines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.Lines(s) {
			if !yield(strings.TrimSpace(line)) {
				return
			}
		}
	}
}

func RawLines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.Lines(s) {
			if !yield(line) {
				return
			}
		}
	}
}
