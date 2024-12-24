package main

import (
	"reflect"
	"testing"
)

func TestGreedy(t *testing.T) {
	actual := encode("abacabad")
	expected := map[rune]string{
		'a': "0",
		'b': "10",
		'c': "110",
		'd': "111",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
