package main

import (
	"reflect"
	"testing"
)

func TestInverse(t *testing.T) {

	ss := []int{0, 7}
	es := []int{5, 10}
	ps := []int{1, 6, 11}
	actual := getPointsRelation(ss, es, ps)

	expeced := []int{1, 0, 0}
	if !reflect.DeepEqual(actual, expeced) {
		t.Errorf("actual = %v, expected = %v", actual, expeced)
	}
}
