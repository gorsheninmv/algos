package main

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {

	xs := []int{2, 3, 9, 2, 9}

	actual := inversions(xs)
	expeced := 2
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}

	xs = []int{5, 4, 3, 2, 1}

	actual = inversions(xs)
	expeced = 10
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}
}
