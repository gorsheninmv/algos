package main

import (
	"reflect"
	"testing"
)

func TestGreedy(t *testing.T) {
	actual := greedy(1)
	expected := []int{1}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = greedy(2)
	expected = []int{2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = greedy(4)
	expected = []int{1, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = greedy(6)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

/**
1:
	1

2:
	2

3:
	1 2

4:
	1 3

5:
	1 4

6:
	1 2 3

7:
	1 2 4

8:
	1 2 5

9:
	1 2 6
**/
