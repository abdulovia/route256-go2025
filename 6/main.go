package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Server struct {
	th  int
	ind int
}

func check(throughputs []Server, weights, ans []int, mt int) (int, bool) {
	res := 0
	for j, weight := range weights {
		curtime := -1
		for i := 0; i < len(throughputs); i++ {
			time := int(math.Ceil(float64(weight)/float64(throughputs[i].th)))
			if time <= mt {
				curtime = time
				ans[j] = throughputs[i].ind + 1
				break
			}
		}
		fmt.Println(res, curtime, weight, mt)
		if curtime == -1 {
			return -1, false
		}
		res = max(res, mt - curtime)
	}

	return res, true
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
		throughputs := make([]Server, n)
		for i := 0; i < n; i++ {
			var throughput int
			fmt.Fscan(in, &throughput)
			throughputs[i] = Server{th: throughput, ind: i}
		}

		// sort.Slice(throughputs, func(i, j int) bool {
		// 	return throughputs[i].th < throughputs[j].th
		// })

		var m int
		fmt.Fscan(in, &m)
		weights := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &weights[i])
		}

		ans := make([]int, m)

		left, right := 0, int(1e9)
		for i := 0; i < 100; i++ {
			mid := (left + right) / 2
			if _, valid := check(throughputs, weights, ans, mid); valid {
				right = mid
			} else {
				left = mid + 1
			}
		}

		result, valid := check(throughputs, weights, ans, left)
		if !valid {
			result, _ = check(throughputs, weights, ans, right)
		}
		fmt.Fprintln(out, result)
		for _, val := range ans {
			fmt.Fprint(out, val, " ")
		}
		fmt.Fprintln(out)
	}
}