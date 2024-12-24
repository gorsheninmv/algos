package main

import (
  "testing"
)

func TestFib(t *testing.T) {
  actual := fib(841645)
  if actual != 5 {
    t.Errorf("actual = %d, expected = 5", actual)
  }
}
