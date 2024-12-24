package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	cost int
	vol  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, capacity int
	fmt.Fscan(in,  &n, &capacity)

	items := make([]Item, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &items[i].cost, &items[i].vol)
	}

	result := greedy(items, capacity)

	fmt.Fprintf(out, "%.3f\n", result)
}

func greedy(items []Item, knapsackVol int) float64 {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].cost*items[j].vol >= items[j].cost*items[i].vol
	})

	vol, price := 0.0, 0.0
	for _, v := range items {
		if vol+float64(v.vol) < float64(knapsackVol) {
			vol = vol + float64(v.vol)
			price = price + float64(v.cost)
		} else {
			// vol + v.vol * k = knapsackVol => k = (knapsackVol - vol) / v.vol
			k := (float64(knapsackVol) - vol) / float64(v.vol)
			price = price + float64(v.cost)*k
			break
		}
	}

	return price
}
