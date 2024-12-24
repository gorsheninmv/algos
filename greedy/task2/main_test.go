package main

import (
	"math"
	"testing"
)

func TestGreedy(t *testing.T) {
	actual := greedy([]Item{{60, 20}, {100, 50}, {120, 30}}, 50)
	expected := 180.0
	if !equalsWithTolerance(actual, expected) {
		t.Errorf("actual = %x, expected = %x", actual, expected)
	}
}

func equalsWithTolerance(a, b float64) bool {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}
