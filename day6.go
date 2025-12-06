package main

import (
	"slices"
	"strconv"
	"strings"
)

type Day6 struct {
	problemChunks [][]string
}

func (d *Day6) Parse(input []byte) {
	lines := slices.Collect(RawLines(string(input)))

	operatorLine := []byte(lines[len(lines)-1])
	var start int
	for start < len(operatorLine) {
		end := slices.IndexFunc(operatorLine[start+1:], func(b byte) bool { return b == '+' || b == '*' })
		if end == -1 {
			end = len(operatorLine) + 1
		} else {
			end += start + 1
		}

		chunk := make([]string, len(lines))
		for i, l := range lines {
			chunk[i] = l[start : end-1]
		}
		d.problemChunks = append(d.problemChunks, chunk)

		start = end
	}
}

func (d *Day6) Part1() int {
	total := 0
	for _, p := range d.problemChunks {
		if p[len(p)-1][0] == '+' {
			subTotal := 0
			for _, l := range p[:len(p)-1] {
				x, _ := strconv.Atoi(strings.TrimSpace(l))
				subTotal += x
			}
			total += subTotal
		} else {
			subTotal := 1
			for _, l := range p[:len(p)-1] {
				x, _ := strconv.Atoi(strings.TrimSpace(l))
				subTotal *= x
			}
			total += subTotal
		}
	}
	return total
}

func (d *Day6) Part2() int {
	total := 0
	for _, p := range d.problemChunks {
		subTotal := 0
		if p[len(p)-1][0] == '*' {
			subTotal = 1
		}

		for col := range len(p[0]) {
			x := 0

			for _, s := range p[:len(p)-1] {
				if s[col] == ' ' {
					continue
				}
				x = (x * 10) + int(s[col]-'0')
			}

			if p[len(p)-1][0] == '+' {
				subTotal += x
			} else {
				subTotal *= x
			}
		}

		total += subTotal
	}

	return total
}
