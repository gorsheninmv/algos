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
	dp := make([]int, len(xs))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if xs[i]%xs[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
