package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	val  int
	next *node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	fmt.Fscan(in, &n)

	xs := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &xs[i])
	}

	result := dp(xs)
	fmt.Fprintln(out, len(result))

	for _, v := range result {
		fmt.Fprintf(out, "%d ", v)
	}
}

func dp(xs []int) []int {
	dp := make([]int, len(xs)+1)
	prev := make([]*node, len(xs)+2)
	size := 1
	for i, v := range xs {
		// binsearch [ >= x ][ < x ]
		left, right := 0, size
		for right-left > 1 {
			mid := left + (right-left)/2
			if dp[mid] >= v {
				left = mid
			} else {
				right = mid
			}
		}
		dp[right] = v
		prev[right+1] = &node{i, prev[right]}
		if right == size {
			size++
		}
	}

	result := make([]int, size-1)
	node := prev[size]
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = node.val + 1
		node = node.next
	}
	return result
}
