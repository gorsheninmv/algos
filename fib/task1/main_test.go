package main

import (
  "testing"
)

func TestFib(t *testing.T) {
  actual := fib(3)
  if actual != 2 {
    t.Errorf("actual = %d, expected = 2", actual)
  }

  actual = fib(1)
  if actual != 1 {
    t.Errorf("actual = %d, expected = 1", actual)
  }
}
