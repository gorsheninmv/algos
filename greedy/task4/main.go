package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NodeHeap []*Node

// #region sort.Interface

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	fmt.Printf("!!! less %d\n", len(h))
	return h[i].freq < h[j].freq
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// #endregion

// #region heap.Interface

// Push and Pop use pointer reveivers because thy modify the slice's length,
// not just its contents.

func (h *NodeHeap) Push(x any) {
	fmt.Println("!!! push")
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// #endregion

type Node struct {
	left  *Node
	right *Node
	char  rune
	freq  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	//result := greedy(n)

}

func encode(s string) map[rune]string {
	nodes := [26]Node{}

	for _, v := range s {
		if nodes[v-'a'].freq == 0 {
			nodes[v-'a'].char = v
		}
		nodes[v-'a'].freq++
	}

	fmt.Printf("!!! %v\n", nodes)

	var h NodeHeap
	for _, node := range nodes {
		if node.freq > 0 {
			heap.Push(&h, &node)
		}
	}

	for len(h) > 1 {
		node1, node2 := heap.Pop(&h).(*Node), heap.Pop(&h).(*Node)
		parent := Node{
			left:  node1,
			right: node2,
			freq:  node1.freq + node2.freq,
		}
		heap.Push(&h, &parent)
	}

	return buildTable(heap.Pop(&h).(*Node))
}

func buildTable(root *Node) map[rune]string {
	fmt.Printf("%v\n", root)
	result := make(map[rune]string)
	var dfs func(node *Node, value int, level int)

	dfs = func(node *Node, value int, level int) {
		if leaf(*node) {
			code := strconv.FormatInt(int64(value), 2)
			prefixLen := level - len(code)
			var sb strings.Builder
			for i := 0; i < prefixLen; i++ {
				sb.WriteRune('0')
			}
			sb.WriteString(code)
			fmt.Println("!!! new")
			result[node.char] = sb.String()
		} else {
			dfs(node.left, value*2, level+1)
			dfs(node.right, value*2+1, level+1)
		}
	}
	dfs(root, 0, 0)
	return result
}

func leaf(node Node) bool {
	return node.left != nil
}
