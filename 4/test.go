package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < 3; i++ {
		b, _ := in.ReadByte()
		if b == '\n' {
			fmt.Fprintln(out, "okay")
		}
		fmt.Fprintln(out, b)
	}
}