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
 
	var a, b int
	fmt.Fscanf(in, "%d %d", &a, &b)

	result := gcd(a, b)
	
	fmt.Fprintf(out, "%d\n", result)
}

// invariant: a >= b
func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}
