package main

import (
  "testing"
)

func TestFib(t *testing.T) {
  actual := fib(10, 2)
  if actual != 1 {
    t.Errorf("actual = %d, expected = 1", actual)
  }
}
