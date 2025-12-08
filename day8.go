package main

import (
	"container/heap"
	"slices"
	"strings"
)

type Day8 struct {
	junctionBoxes []Point3D

	junctionBoxPairHeap *pointPairDistanceHeap
}

func (d *Day8) Parse(input []byte) {
	for l := range Lines(string(input)) {
		parts := ParseInts(strings.Split(l, ","))
		d.junctionBoxes = append(d.junctionBoxes, Point3D{parts[0], parts[1], parts[2]})
	}

	_pairHeap := make(pointPairDistanceHeap, 0, len(d.junctionBoxes)*(len(d.junctionBoxes)-1)/2)
	d.junctionBoxPairHeap = &_pairHeap
	for i, p1 := range d.junctionBoxes {
		for j := i + 1; j < len(d.junctionBoxes); j++ {
			p2 := d.junctionBoxes[j]
			dist := p1.SquaredEuclideanDist(p2)
			*d.junctionBoxPairHeap = append(*d.junctionBoxPairHeap, struct {
				a Point3D
				b Point3D
				d int
			}{p1, p2, dist})
		}
	}
	heap.Init(d.junctionBoxPairHeap)
}

func (d *Day8) Part1() int {
	_pairHeap := slices.Clone(*d.junctionBoxPairHeap)
	pairHeap := &_pairHeap

	circuits := []Set[Point3D]{}
	pointToCircuit := make(map[Point3D]int)

	for range 1000 {
		pair := heap.Pop(pairHeap).(struct {
			a, b Point3D
			d    int
		})

		ca, oka := pointToCircuit[pair.a]
		cb, okb := pointToCircuit[pair.b]

		if oka && okb {
			if ca == cb {
				continue
			}

			// Move all in cb to ca
			for p := range circuits[cb] {
				pointToCircuit[p] = ca
				circuits[ca].Insert(p)
			}
			circuits[cb] = nil
			continue
		} else if oka && !okb {
			pointToCircuit[pair.b] = ca
			circuits[ca].Insert(pair.b)
		} else if !oka && okb {
			pointToCircuit[pair.a] = cb
			circuits[cb].Insert(pair.a)
		} else {
			c := len(circuits)
			circuits = append(circuits, NewSet(pair.a, pair.b))
			pointToCircuit[pair.a] = c
			pointToCircuit[pair.b] = c
		}
	}

	largest := [3]Set[Point3D]{nil, nil, nil}
	for _, c := range circuits {
		if len(c) > len(largest[0]) {
			largest = [3]Set[Point3D]{c, largest[0], largest[1]}
		} else if len(c) > len(largest[1]) {
			largest = [3]Set[Point3D]{largest[0], c, largest[1]}
		} else if len(c) > len(largest[2]) {
			largest = [3]Set[Point3D]{largest[0], largest[1], c}
		}
	}

	return len(largest[0]) * len(largest[1]) * len(largest[2])
}

func (d *Day8) Part2() int {
	validCircuits := 0
	circuits := []Set[Point3D]{}
	pointToCircuit := make(map[Point3D]int)

	for {
		pair := heap.Pop(d.junctionBoxPairHeap).(struct {
			a, b Point3D
			d    int
		})

		ca, oka := pointToCircuit[pair.a]
		cb, okb := pointToCircuit[pair.b]

		if oka && okb {
			if ca == cb {
				continue
			}

			// Move all in cb to ca
			for p := range circuits[cb] {
				pointToCircuit[p] = ca
				circuits[ca].Insert(p)
			}
			circuits[cb] = nil

			validCircuits--
			if validCircuits == 1 && len(pointToCircuit) == len(d.junctionBoxes) {
				return pair.a.X * pair.b.X
			}

			continue
		} else if oka && !okb {
			pointToCircuit[pair.b] = ca
			circuits[ca].Insert(pair.b)
		} else if !oka && okb {
			pointToCircuit[pair.a] = cb
			circuits[cb].Insert(pair.a)
		} else {
			c := len(circuits)
			circuits = append(circuits, NewSet(pair.a, pair.b))
			pointToCircuit[pair.a] = c
			pointToCircuit[pair.b] = c
			validCircuits++
		}
	}
}

type pointPairDistanceHeap []struct {
	a, b Point3D
	d    int
}

var _ heap.Interface = &pointPairDistanceHeap{}

func (h *pointPairDistanceHeap) Len() int {
	return len(*h)
}

func (h *pointPairDistanceHeap) Less(i, j int) bool {
	return (*h)[i].d < (*h)[j].d
}

func (h *pointPairDistanceHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *pointPairDistanceHeap) Push(x any) {
	*h = append(*h, x.(struct {
		a, b Point3D
		d    int
	}))
}

func (h *pointPairDistanceHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
