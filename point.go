package main

type Point struct {
	X int
	Y int
}

func (p Point) AdjacentPoints() [8]Point {
	return [8]Point{
		{p.X - 1, p.Y - 1}, {p.X, p.Y - 1}, {p.X + 1, p.Y - 1},
		{p.X - 1, p.Y} /*               */, {p.X + 1, p.Y},
		{p.X - 1, p.Y + 1}, {p.X, p.Y + 1}, {p.X + 1, p.Y + 1},
	}
}
