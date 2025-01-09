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

	result := dp(xs)
	fmt.Fprintln(out, result)
}

func dp(xs []int) int {
	dp := make([]int, len(xs)+1)
	size := 1
	for _, v := range xs {
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
		if right == size {
			size++
		}
	}
	fmt.Printf("%v\n", dp)
	return size - 1
}
