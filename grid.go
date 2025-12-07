package main

import (
	"bytes"
	"iter"
	"slices"
)

type Grid struct {
	elements []byte
	rows     int
	cols     int
}

func NewGrid(input []byte) Grid {
	rows := bytes.Split(input, []byte("\n"))
	cols := len(rows[0])
	return Grid{
		elements: slices.Concat(rows...),
		rows:     len(rows),
		cols:     cols,
	}
}

func (g Grid) Len() int {
	return len(g.elements)
}

func (g Grid) Valid(p Point) bool {
	return p.X >= 0 && p.X < g.cols &&
		p.Y >= 0 && p.Y < g.rows
}

func (g Grid) At(p Point) (byte, bool) {
	if !g.Valid(p) {
		return 0, false
	}
	return g.elements[p.Y*g.cols+p.X], true
}

func (g Grid) Set(p Point, b byte) bool {
	if !g.Valid(p) {
		return false
	}
	g.elements[p.Y*g.cols+p.X] = b
	return true
}

func (g Grid) All() iter.Seq2[byte, Point] {
	return func(yield func(byte, Point) bool) {
		for y := 0; y < g.rows; y++ {
			for x := 0; x < g.cols; x++ {
				b := g.elements[y*g.cols+x]
				if !yield(b, Point{x, y}) {
					return
				}
			}
		}
	}
}

func (g Grid) Find(target byte) (Point, bool) {
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			b := g.elements[y*g.cols+x]
			if b == target {
				return Point{x, y}, true
			}
		}
	}

	return Point{}, false
}

func (g Grid) FindAll(target byte) iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for y := 0; y < g.rows; y++ {
			for x := 0; x < g.cols; x++ {
				b := g.elements[y*g.cols+x]
				if b == target {
					if !yield(Point{x, y}) {
						return
					}
				}
			}
		}
	}
}
