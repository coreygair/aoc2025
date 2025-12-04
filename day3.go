package main

import "math"

type Day3 struct {
	banks [][]int
}

func (d *Day3) Parse(input []byte) {
	for l := range Lines(string(input)) {
		bank := make([]int, len(l))
		for i, c := range l {
			bank[i] = int(c - '0')
		}
		d.banks = append(d.banks, bank)
	}
}

func (d *Day3) Part1() int {
	totalJoltage := 0

	for _, bank := range d.banks {
		var max, maxI int
		for i, x := range bank[:len(bank)-1] {
			if x > max {
				max = x
				maxI = i
			}
		}

		var nextMax int
		for _, x := range bank[maxI+1:] {
			if x > nextMax {
				nextMax = x
			}
		}

		joltage := (10 * max) + nextMax
		totalJoltage += joltage
	}

	return totalJoltage
}

func (d *Day3) Part2() int {
	totalJoltage := 0

	for _, bank := range d.banks {
		joltage := 0
		pos := 0
		for digitsLeft := 11; digitsLeft >= 0; digitsLeft-- {
			var max, maxI int
			for i, x := range bank[pos : len(bank)-digitsLeft] {
				if x > max {
					max = x
					maxI = i
				}
			}

			joltage += max * int(math.Pow10(digitsLeft))
			pos += maxI + 1
		}

		totalJoltage += joltage
	}

	return totalJoltage
}
