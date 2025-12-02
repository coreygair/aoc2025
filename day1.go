package main

import (
	"strconv"
)

type Day1 struct {
	instructions []int
}

func (d *Day1) Parse(input []byte) {
	for line := range Lines(string(input)) {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}

		n, _ := strconv.Atoi(line[1:])
		d.instructions = append(d.instructions, n*sign)
	}
}

func (d *Day1) Part1() int {
	numZero := 0

	dial := 50
	for _, n := range d.instructions {
		dial = (dial + n) % 100
		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			numZero++
		}
	}

	return numZero
}

func (d *Day1) Part2() int {
	numThroughZero := 0

	dial := 50
	for _, n := range d.instructions {
		numFullRotations := n / 100
		partialRoatation := n - (100 * numFullRotations)

		if numFullRotations < 0 {
			numFullRotations = -numFullRotations
		}
		numThroughZero += numFullRotations

		if dial != 0 {
			if (n < 0 && dial <= -partialRoatation) || (n > 0 && 100-dial <= partialRoatation) {
				numThroughZero++
			}
		}

		dial = (dial + n) % 100
		if dial < 0 {
			dial += 100
		}
	}

	return numThroughZero
}
