package main

import (
  "testing"
)

func TestGcd(t *testing.T) {
  actual := gcd(18, 35)
  if actual != 1 {
    t.Errorf("actual = %d, expected = 1", actual)
  }

  actual = gcd(14159572, 63967072)
  if actual != 4 {
    t.Errorf("actual = %d, expected = 4", actual)
  }
}
