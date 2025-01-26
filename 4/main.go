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

	var t int
	fmt.Fscan(in, &t)

	for inp := 0; inp < t; inp++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadByte()

		evenpat := make(map[string]int)
		oddpat := make(map[string]int)
		strs := make(map[string]int)

		for i := 0; i < n; i++ {
			var even []byte
			var odd []byte
			var str []byte

			for j := 0; ; j++ {
				c, _ := in.ReadByte()
				if c == '\n' {
					break
				}
				if j%2 == 0 {
					even = append(even, c)
				} else {
					odd = append(odd, c)
				}
				str = append(str, c)
			}

			s := string(str)
			evenpat[string(even)]++
			if len(s) > 1 {
				oddpat[string(odd)]++
			}
			strs[s]++
		}

		count := 0
		for _, N := range evenpat {
			if N > 1 {
				count += N * (N - 1) / 2
			}
		}
		for _, N := range oddpat {
			if N > 1 {
				count += N * (N - 1) / 2
			}
		}
		for st, N := range strs {
			if len(st) > 1 && N > 1 {
				count -= N * (N - 1) / 2
			}
		}

		fmt.Fprintln(out, count)
	}
}
