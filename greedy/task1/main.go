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

	var n int
	fmt.Fscanf(in, "%d", &n)

	segments := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &segments[i][0], &segments[i][1])
	}

	result := greedy(segments)

	fmt.Fprintf(out, "%d\n", len(result))
	for i := 0; i < len(result); i++ {
		fmt.Fprintf(out, "%d ", result[i])
	}
}

func greedy(segments [][2]int) []int {
	sort.SliceStable(segments, func(i, j int) bool {
		return segments[i][1] < segments[j][1]
	})

	point := -1
	var result []int
	for i := 0; i < len(segments); i++ {
		if point < segments[i][0] {
			point = segments[i][1]
			result = append(result, point)
		}
	}

	return result
}
