package main

import (
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	freshRanges [][]int
	available   []int
}

func (d *Day5) Parse(input []byte) {
	lines := slices.Collect(Lines(string(input)))

	mid := slices.Index(lines, "")

	for _, l := range lines[:mid] {
		parts := strings.Split(l, "-")
		lo, _ := strconv.Atoi(parts[0])
		hi, _ := strconv.Atoi(parts[1])

		d.freshRanges = append(d.freshRanges, []int{lo, hi})
	}

	for _, l := range lines[mid+1:] {
		x, _ := strconv.Atoi(l)
		d.available = append(d.available, x)
	}
}

func (d *Day5) Part1() int {
	freshCount := 0
	for _, x := range d.available {
		for _, r := range d.freshRanges {
			if x >= r[0] && x <= r[1] {
				freshCount++
				break
			}
		}
	}
	return freshCount
}

func (d *Day5) Part2() int {
	rangeStarts := make([]int, 0, len(d.freshRanges))
	rangeEnds := make([]int, 0, len(d.freshRanges))
	for _, r := range d.freshRanges {
		rangeStarts = append(rangeStarts, r[0])
		rangeEnds = append(rangeEnds, r[1])
	}
	slices.Sort(rangeStarts)
	slices.Sort(rangeEnds)

	totalFreshIDs := 0

	countInRanges := 0
	rangeStart := 0
	i, j := 0, 0
	for j < len(rangeEnds) {
		if i < len(rangeStarts) && rangeStarts[i] <= rangeEnds[j] {
			if countInRanges == 0 {
				rangeStart = rangeStarts[i]
			}
			countInRanges++
			i++
		} else {
			countInRanges--
			if countInRanges == 0 {
				totalFreshIDs += (rangeEnds[j] - rangeStart) + 1
			}
			j++
		}
	}

	return totalFreshIDs
}
