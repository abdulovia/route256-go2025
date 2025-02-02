package main

import (
	"bufio"
	"fmt"
	"os"
)

type Answer struct {
	x, y int
	dir  string
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for inp := 0; inp < t; inp++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		
		ans := []Answer{}

		if n < m {
			if n == 1 {
				ans = append(ans, Answer{1, 1, "R"})
			} else {
				ans = append(ans, Answer{1, 1, "R"})
				ans = append(ans, Answer{n, m, "L"})
			}
		} else {
			if m == 1 {
				ans = append(ans, Answer{1, 1, "D"})
			} else {
				ans = append(ans, Answer{1, 1, "D"})
				ans = append(ans, Answer{n, m, "U"})
			}
		}

		fmt.Fprintln(out, len(ans))
		for _, an := range ans {
			fmt.Fprintln(out, an.x, an.y, an.dir)
		}
	}
}
