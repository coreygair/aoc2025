package main

import "strconv"

func ParseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func ParseInts(s []string) []int {
	result := make([]int, len(s))
	for i, s := range s {
		x, _ := strconv.Atoi(s)
		result[i] = x
	}
	return result
}
