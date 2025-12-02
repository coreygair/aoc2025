package main

import (
	"math"
	"strconv"
	"strings"
)

type Day2 struct {
	rangesToCheck []struct{ first, last int }
}

func (d *Day2) Parse(input []byte) {
	parts := strings.Split(string(input), ",")
	d.rangesToCheck = make([]struct {
		first int
		last  int
	}, len(parts))

	for i, part := range parts {
		parts := strings.Split(part, "-")
		first, _ := strconv.Atoi(parts[0])
		last, _ := strconv.Atoi(parts[1])

		d.rangesToCheck[i] = struct {
			first int
			last  int
		}{first, last}
	}
}

func (d *Day2) Part1() int {
	sumOfInvalidIDs := 0

	for _, r := range d.rangesToCheck {
		for id := r.first; id <= r.last; id++ {
			digits := int(math.Log10(float64(id))) + 1

			if digits%2 == 1 {
				continue
			}

			halfDigitsShift := int(math.Pow10(digits / 2))
			upperPart := id / halfDigitsShift
			lowerPart := id - (upperPart * halfDigitsShift)

			if upperPart == lowerPart {
				sumOfInvalidIDs += id
			}
		}
	}

	return sumOfInvalidIDs
}

func (d *Day2) Part2() int {
	sumOfInvalidIDs := 0

	for _, r := range d.rangesToCheck {
		for id := r.first; id <= r.last; id++ {
			digits := int(math.Log10(float64(id))) + 1

			for subseqDigits := 1; subseqDigits <= digits/2; subseqDigits++ {
				if digits%subseqDigits != 0 {
					continue
				}

				subseq := id / int(math.Pow10(digits-subseqDigits))

				reps := digits / subseqDigits
				subseqDigitsShift := int(math.Pow10(subseqDigits))

				test := 0
				for j := range reps {
					test += subseq * int(math.Pow(float64(subseqDigitsShift), float64(j)))
				}

				if test == id {
					sumOfInvalidIDs += id
					break
				}
			}
		}
	}

	return sumOfInvalidIDs
}
