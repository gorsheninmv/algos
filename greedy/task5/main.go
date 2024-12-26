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

	var letters, length int
	fmt.Fscanf(in, "%d %d\n", &letters, &length)

	codemap := make(map[rune]string)
	for i := 0; i < 4; i++ {
		var ch rune
		var code string
		fmt.Fscanf(in, "%c: %s\n", &ch, &code)
		codemap[ch] = code
	}

	fmt.Printf("%v\n", codemap)
}
