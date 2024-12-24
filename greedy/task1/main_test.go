package main

import (
	"reflect"
	"testing"
)

func TestGreedy(t *testing.T) {
	actual := greedy([][2]int{{1, 3}, {2, 5}, {3, 6}})
	expected := []int{3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %x, expected = %x", actual, expected)
	}

	actual = greedy([][2]int{{4, 7}, {1, 3}, {2, 5}, {5, 6}})
	expected = []int{3, 6}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %x, expected = %x", actual, expected)
	}
}
