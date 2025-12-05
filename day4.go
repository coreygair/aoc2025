package main

const (
	empty byte = '.'
	paper byte = '@'
)

type Day4 struct {
	grid Grid
}

func (d *Day4) Parse(input []byte) {
	d.grid = NewGrid(input)
}

func (d *Day4) Part1() int {
	numAccessible := 0
	for p := range d.grid.FindAll(paper) {
		if d.countAdjacentPaper(p) < 4 {
			numAccessible++
		}
	}
	return numAccessible
}

func (d *Day4) Part2() int {
	adjacentPaperCounts := make(map[Point]int, d.grid.Len())
	for p := range d.grid.FindAll(paper) {
		adjacentPaperCounts[p] = d.countAdjacentPaper(p)
	}

	removed := 0

	for {
		toRemove := make(Set[Point])
		for p, c := range adjacentPaperCounts {
			if c < 4 {
				toRemove.Insert(p)
				delete(adjacentPaperCounts, p)
			}
		}

		if len(toRemove) == 0 {
			break
		}

		removed += len(toRemove)

		for p := range toRemove {
			for _, p := range p.AdjacentPoints() {
				old, ok := adjacentPaperCounts[p]
				if !ok {
					continue
				}

				adjacentPaperCounts[p] = old - 1
			}
		}
	}

	return removed
}

func (d *Day4) countAdjacentPaper(p Point) int {
	paperAdjacent := 0
	for _, p := range p.AdjacentPoints() {
		if b, ok := d.grid.At(p); ok && b == paper {
			paperAdjacent++
		}
	}
	return paperAdjacent
}
