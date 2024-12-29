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
	fmt.Fscanf(in, "%d\n", &n)

	var ch rune
	var val int
	var heap heap
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%c", &ch)
		if ch == 'I' {
			in.Discard(6)
			fmt.Fscanf(in, "%d\n", &val)
			heap.insert(val)
		} else {
			in.ReadBytes('\n')
			fmt.Fprintln(out, heap.extractmax())
		}
	}
}

func (heap *heap) insert(v int) {
	*heap = append(*heap, v)
	heap.siftup(len(*heap) - 1)
}

func (h *heap) extractmax() int {
	heap := *h
	ret := heap[0]
	heap[0] = heap[len(heap)-1]
	*h = heap[:len(heap)-1]
	heap.siftdown(0)
	return ret
}

func (h *heap) siftup(idx int) {
	if idx == 0 {
		return
	}

	parentIdx := (idx - 1) / 2
	heap := *h
	if heap[parentIdx] < heap[idx] {
		heap[parentIdx], heap[idx] = heap[idx], heap[parentIdx]
		h.siftup(parentIdx)
	}
}

func (h *heap) siftdown(idx int) {
	heap := *h
	if heap.isleaf(idx) {
		return
	}

	maxIdx := idx
	if leftIdx := idx*2 + 1; heap[maxIdx] < heap[leftIdx] {
		maxIdx = leftIdx
	}
	if rightIdx := idx*2 + 2; rightIdx < len(heap) && heap[maxIdx] < heap[rightIdx] {
		maxIdx = rightIdx
	}

	if maxIdx != idx {
		heap[idx], heap[maxIdx] = heap[maxIdx], heap[idx]
		h.siftdown(maxIdx)
	}
}

func (heap heap) isleaf(idx int) bool {
	leftIdx := idx*2 + 1
	return leftIdx >= len(heap)
}
