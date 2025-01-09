package main

import (
	"bufio"
	"fmt"
	"os"
)

type heap []int

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

	fmt.Fscan(in, &n)
	ts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ts[i])
	}

	for _, t := range ts {
		x := binsearch(xs, t)
		if x == -1 {
			fmt.Fprintf(out, "%d ", -1)
		} else {
			fmt.Fprintf(out, "%d ", x+1)
		}
	}
}

func binsearch(xs []int, target int) int {
	// [ <= target ] [ > target ]
	//             l r

	l, r := -1, len(xs)
	for r-l > 1 {
		m := l + (r-l)/2
		if xs[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	if l >= 0 && xs[l] == target {
		return l
	} else {
		return -1
	}
}
