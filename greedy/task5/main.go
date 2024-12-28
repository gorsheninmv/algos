package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var letters, length int
	fmt.Fscanf(in, "%d %d\n", &letters, &length)

	codemap := make(map[string]rune)
	for i := 0; i < letters; i++ {
		var ch rune
		var code string
		fmt.Fscanf(in, "%c: %s\n", &ch, &code)
		codemap[code] = ch
	}

	var code string
	fmt.Fscanf(in, "%s\n", &code)
	decode(in, out, codemap)
}

func decode(reader io.RuneReader, writer io.ByteWriter, codemap map[string]rune) {
	buf := ""
	for {
		if v, _, err := reader.ReadRune(); err == nil {
			buf = buf + string(v)
			if letter, ok := codemap[buf]; ok {
				writer.WriteByte(byte(letter))
				buf = ""
			}
		} else {
			break
		}
	}
}
