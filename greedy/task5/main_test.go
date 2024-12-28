package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	x := strings.NewReader("abcdef")
	var y bytes.Buffer
	decode(x, &y, make(map[string]rune)) 
	expected := map[rune]string{
		'a': "0",
		'b': "10",
		'c': "110",
		'd': "111",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	actual = codemap(huffmantree("a"))
	expected = map[rune]string{
		'a': "0",
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

	actual = codelen(huffmantree("a"))
	expected = 1
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
