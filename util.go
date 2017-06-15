package gago

import (
	"math"
)

func newInts(n int) []int {
	var ints = make([]int, n)
	for i := range ints {
		ints[i] = i
	}
	return ints
}

// Divide each element in a float64 slice by a given value.
func divide(floats []float64, value float64) []float64 {
	var divided = make([]float64, len(floats))
	for i, v := range floats {
		divided[i] = v / value
	}
	return divided
}

// Compute the cumulative sum of a float64 slice.
func cumsum(floats []float64) []float64 {
	var summed = make([]float64, len(floats))
	copy(summed, floats)
	for i := 1; i < len(summed); i++ {
		summed[i] += summed[i-1]
	}
	return summed
}

// Find the minimum between two ints.
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Compute the sum of a float64 slice.
func sumFloat64s(floats []float64) (sum float64) {
	for _, v := range floats {
		sum += v
	}
	return
}

// Compute the minimum value of a float64 slice.
func minFloat64s(floats []float64) (min float64) {
	min = math.Inf(1)
	for _, f := range floats {
		if f < min {
			min = f
		}
	}
	return
}

// Compute the maximum value of a float64 slice.
func maxFloat64s(floats []float64) (max float64) {
	max = math.Inf(-1)
	for _, f := range floats {
		if f > max {
			max = f
		}
	}
	return
}

// Compute the mean of a float64 slice.
func meanFloat64s(floats []float64) float64 {
	return sumFloat64s(floats) / float64(len(floats))
}

// Compute the variance of a float64 slice.
func varianceFloat64s(floats []float64) float64 {
	var squares = make([]float64, len(floats))
	for i, f := range floats {
		squares[i] = math.Pow(f, 2)
	}
	return meanFloat64s(squares) - math.Pow(meanFloat64s(floats), 2)
}

type set map[interface{}]bool

// union merges two slices and ignores duplicates.
func union(x, y set) set {
	var (
		u         = make(set)
		blackList = make(map[interface{}]bool)
	)
	for i := range x {
		u[i] = true
		blackList[i] = true
	}
	for i := range y {
		if !blackList[i] {
			u[i] = true
			blackList[i] = true
		}
	}
	return u
}
