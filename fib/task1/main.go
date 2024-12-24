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
	fmt.Fscanf(in, "%d", &n)

	result := fib(n)
	
	fmt.Fprintf(out, "%d\n", result)
}

func fib(n int) int {
	n1, n2 := 0, 1
	for i := 2; i <= n; i++ {
		next := n1 + n2
		n1 = n2
		n2 = next
	}
	return n2
}
