package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNotValidName(name string) bool {
	for _, l := range name {
		if !(l >= 'a' && l <= 'z') {
			return true
		}
	}
	return false
}

func isValid(items map[string]int, line string) bool {
	prcs := make(map[int]bool)
	itms := make(map[string]int)
	if strings.Index(line, " ") != -1 {
		return false
	}
	els := strings.Split(line, ",")

	for _, el := range els {
		l := strings.Split(el, ":")
		if len(l) != 2 {
			return false
		}
		if len(l[0]) > 10 || len(l[1]) > 10 || isNotValidName(l[0]) {
			return false
		}
		name := l[0]
		price, err := strconv.Atoi(l[1])
		if err != nil || price < 1 || price > int(1e9) || l[1][0] == '0' {
			return false
		}

		if _, ok := prcs[price]; ok {
			return false
		}
		if _, ok := itms[name]; ok {
			return false
		}
		if val, ok := items[name]; !ok || val != price {
			return false
		}

		prcs[price] = true
		itms[name] = price
	}

	for _, price := range items {
		if _, ok := prcs[price]; !ok {
			return false
		}
	}

	return true
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
		var n int
		fmt.Fscan(in, &n)

		items := make(map[string]int)
		for i := 0; i < n; i++ {
			var name string
			var price int
			fmt.Fscan(in, &name, &price)
			items[name] = price
		}

		var line string
		fmt.Fscan(in, &line)

		if isValid(items, line) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
