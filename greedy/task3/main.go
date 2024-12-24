package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	cost int
	vol  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	result := greedy(n)

	fmt.Println(len(result))
	for _, v := range result {
		fmt.Fprintf(out, "%d ", v)
	}
}

func greedy(n int) []int {
	var result []int
	sum := 0
	for i := 1; i <= n; i++ {
		if (sum + i) < n {
			result = append(result, i)
			sum = sum + i
		} else if (sum + i) == n {
			result = append(result, i)
			break
		} else {
			result[len(result)-1] = n - sum + result[len(result)-1]
			break
		}
	}
	return result
}
