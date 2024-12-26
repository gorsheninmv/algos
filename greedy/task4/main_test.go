package main

import (
	"reflect"
	"testing"
)

func TestCodemap(t *testing.T) {
	actual := codemap(huffmantree("abacabad"))
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

func TestCodelen(t *testing.T) {
	actual := codelen(huffmantree("abacabad"))
	expected := 14
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
