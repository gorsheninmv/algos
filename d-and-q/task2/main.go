package main

import (
	"bufio"
	"fmt"
	"os"
)

// For each element in the left, increment a counter based on the number of elements in the right

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xs[i])
	}

	x := inversions(xs)
	fmt.Fprintln(out, x)
}

func inversions(xs []int) int {
	var inversions int
	var merge func(left []int, right []int) []int
	var mergesort func(xs []int) []int

	mergesort = func(xs []int) []int {
		if len(xs) == 1 {
			return xs
		} else {
			m := len(xs) / 2
			return merge(mergesort(xs[:m]), mergesort(xs[m:]))
		}
	}

	merge = func(left, right []int) []int {
		i, j := 0, 0

		result := make([]int, len(left)+len(right))
		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				result[i+j] = left[i]
				i++
			} else {
				result[i+j] = right[j]
				inversions = inversions + (len(left) - i)
				j++
			}
		}

		for i < len(left) {
			result[i+j] = left[i]
			i++
		}

		for j < len(right) {
			result[i+j] = right[j]
			j++
		}

		return result
	}

	mergesort(xs)
	return inversions
}
