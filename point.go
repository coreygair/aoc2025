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

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) SquaredEuclideanDist(other Point3D) int {
	dx := Abs(p.X - other.X)
	dy := Abs(p.Y - other.Y)
	dz := Abs(p.Z - other.Z)
	return dx*dx + dy*dy + dz*dz
}

func Abs[T ~int | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
