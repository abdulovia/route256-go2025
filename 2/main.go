package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		var n string
		fmt.Fscan(in, &n)

		x, _ := strconv.Atoi(n)

		if x < 10 {
			fmt.Fprintln(out, x+1)
			continue
		}

		ans := (len(n)-1)*10
		fd, _ := strconv.Atoi(string(n[0]))

		mxn, _ := strconv.Atoi(strings.Repeat(string(n[0]), len(n)))
		ans += fd - 1
		if x >= mxn {
			ans += 1
		}
		fmt.Fprintln(out, ans)
	}
}
