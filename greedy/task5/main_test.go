package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	encoded := strings.NewReader("01001100100111")
	var buf bytes.Buffer
	codemap := map[string]rune{
		"0": 'a',
		"10": 'b',
		"110": 'c',
		"111": 'd',
	}
	decode(encoded, &buf, codemap) 
	actual := buf.String()
	expected := "abacabad"
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}

	encoded = strings.NewReader("0")
	codemap = map[string]rune{
		"0": 'a',
	}
	buf.Reset()
	decode(encoded, &buf, codemap) 
	actual = buf.String()
	expected = "a"
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}
