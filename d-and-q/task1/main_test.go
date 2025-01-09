package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {

	xs := []int{1, 5, 8, 12, 13}
	ts := []int{8, 1, 23, 1, 11}
	es := []int{2, 0, -1, 0, -1}

	for i := 0; i < len(ts); i++ {
		actual := binsearch(xs, ts[i])
		if !reflect.DeepEqual(actual, es[i]) {
			t.Errorf("actual = %v, expected = %v", actual, es[i])
		}
	}
}
