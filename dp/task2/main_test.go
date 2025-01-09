package main

import (
	"reflect"
	"testing"
)

func TestDP(t *testing.T) {
	xs := []int{5, 3, 4, 4, 2}
	actual := dp(xs)

	expeced := 4
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}

	xs = []int{3, 3, 3}
	actual = dp(xs)

	expeced = 3
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}

	xs = []int{1}
	actual = dp(xs)

	expeced = 1
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}
}
