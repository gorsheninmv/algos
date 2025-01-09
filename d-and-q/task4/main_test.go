package main

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {

	xs := []int{2, 3, 9, 2, 9}
	actual := countSort(xs)

	expeced := []int{2, 2, 3, 9, 9}
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}
}
