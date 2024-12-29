package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	var heap heap
	heap.insert(200)
	heap.insert(10)
	heap.insert(5)
	heap.insert(500)

	actual := heap.extractmax()
	expected := 500
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = heap.extractmax()
	expected = 200
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = heap.extractmax()
	expected = 10
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = heap.extractmax()
	expected = 5
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
