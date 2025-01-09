package main

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {

	xs := []int{3, 6, 7, 12}
	actual := dp(xs)

	expeced := 3
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}

	xs = []int{3, 3, 3}
	actual = dp(xs)

	expeced = 3
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}

	xs = []int{1, 4, 4, 6, 7, 9, 12, 14, 14, 15, 17, 20, 22, 24, 24, 27, 30, 30}
	actual = dp(xs)

	expeced = 6
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}
}
