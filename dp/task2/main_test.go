package main

import (
	"reflect"
	"testing"
)

func TestDP(t *testing.T) {
	xs := []int{5, 3, 4, 4, 2}
	actual := dp(xs)

	expected := []int{1, 3, 4, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	xs = []int{3, 3, 3}
	actual = dp(xs)

	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	xs = []int{1}
	actual = dp(xs)

	expected = []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	xs = []int{3, 2, 1, 10, 9, 8, 7}
	actual = dp(xs)

	expected = []int{4, 5, 6, 7}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
