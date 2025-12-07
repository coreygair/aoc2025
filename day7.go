package main

type Day7 struct {
	grid  Grid
	start Point
}

func (d *Day7) Parse(input []byte) {
	d.grid = NewGrid(input)
	s, _ := d.grid.Find('S')
	d.start = s
}

func (d *Day7) Part1() int {
	splittersHit := make(Set[Point])

	toCheck := NewSet(d.start)
	for len(toCheck) > 0 {
		newToCheck := make(Set[Point])

		for p := range toCheck {
			p.Y += 1

			b, ok := d.grid.At(p)
			if !ok {
				continue
			}

			if b == '^' {
				splittersHit.Insert(p)
				newToCheck.Insert(Point{p.X - 1, p.Y}, Point{p.X + 1, p.Y})
				continue
			}

			newToCheck.Insert(p)
		}

		toCheck = newToCheck
	}

	return len(splittersHit)
}

func (d *Day7) Part2() int {

	toCheck := map[Point]int{d.start: 1}
outer:
	for {
		newToCheck := make(map[Point]int)

		for p, timelines := range toCheck {
			p.Y += 1

			b, ok := d.grid.At(p)
			if !ok {
				// As we scan rows sequentially and the splitters never put a beam off the sides of the grid (by inspection of the input)
				// the first time we see an out of bounds point we have reached the bottom of the grid and need to count the timelines.
				break outer
			}

			if b == '^' {
				newToCheck[Point{p.X - 1, p.Y}] += timelines
				newToCheck[Point{p.X + 1, p.Y}] += timelines
				continue
			}

			newToCheck[p] += timelines
		}

		toCheck = newToCheck
	}

	totalTimelines := 0
	for _, t := range toCheck {
		totalTimelines += t
	}
	return totalTimelines
}
