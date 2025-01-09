package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int

	fmt.Fscan(in, &n, &m)
	ss := make([]int, n)
	es := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ss[i], &es[i])
	}

	ps := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &ps[i])
	}

	result := getPointsRelation(ss, es, ps)
	for _, v := range result {
		fmt.Fprintf(out, "%d ", v)
	}
}

func getPointsRelation(ss []int, es []int, ps []int) []int {
	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	sort.SliceStable(es, func(i, j int) bool {
		return es[i] < es[j]
	})
	result := make([]int, len(ps))
	for i, v := range ps {
		segmentStartPointIn := binsearch(ss, v) + 1     // число начал слева, учитывая точку
		segmentFinishPointOut := binsearch(es, v-1) + 1 // число концов слева, не учитывая точку
		pc := segmentStartPointIn - segmentFinishPointOut
		result[i] = pc
	}
	return result
}

func binsearch(xs []int, target int) int {
	// [ <= target ][ > target ]
	l, r := -1, len(xs)
	for r-l > 1 {
		m := l + (r-l)/2
		if xs[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	return l
}
