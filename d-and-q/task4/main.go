package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	xs = countSort(xs)
	for _, v := range xs {
		fmt.Fprintf(out, "%d ", v)
	}
}

func countSort(xs []int) []int {
	counters := make([]int, 11)
	for _, v := range xs {
		counters[v]++
	}

	for i := 1; i < len(counters); i++ {
		counters[i] = counters[i-1] + counters[i]
	}

	result := make([]int, len(xs))
	for i := len(xs) - 1; i >= 0; i-- {
		counters[xs[i]]--
		result[counters[xs[i]]] = xs[i]
	}

	return result
}
