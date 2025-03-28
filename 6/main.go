package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Data struct {
	time int
	wtI  int
	thI  int
}

type Datas []*Data

func (s Datas) Len() int      { return len(s) }
func (s Datas) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByTime struct{ Datas }

func (s ByTime) Less(i, j int) bool { return s.Datas[i].time < s.Datas[j].time }

func twop(allt *Datas, ans *[]int, n int, exit bool, ans2 int) int {
	an := int(2e9)
	wcnt := make([]int, n) // the number of weight in window allt[left, right]
	cnt := 0
	left := 0
	right := 0
outerLoop:
	for right < len(*allt) {
		for cnt != n && right < len(*allt) {
			dat := (*allt)[right]
			(*ans)[dat.wtI] = dat.thI + 1
			wcnt[dat.wtI] += 1
			if wcnt[dat.wtI] == 1 {
				cnt += 1
			}
			if cnt == n {
				break
			}
			right += 1
		}
		for {
			if cnt != n {
				break
			}
			dat := (*allt)[left]
			wcnt[dat.wtI] -= 1
			if wcnt[dat.wtI] == 0 {
				diff := (*allt)[right].time - (*allt)[left].time
				if exit && diff == ans2 {
					break outerLoop
				}
				an = min(an, diff)
				cnt -= 1
				left += 1
				right += 1
				break
			}
			left += 1
		}
	}
	return an
}

func solve(out *bufio.Writer, th, wt *[]int) {
	var allt Datas
	for wtI, w := range *wt {
		for thI, v := range *th {
			time := (w - 1) / v
			allt = append(allt, &Data{time: time, wtI: wtI, thI: thI})
		}
	}
	sort.Sort(ByTime{allt})
	// fmt.Fprint(out, allt[0].time, allt[1].time, allt[2].time)
	ans := make([]int, len(*wt)) // the server assigned to weight: ans[wtI] = thI
	ans1 := twop(&allt, &ans, len(*wt), false, -1)
	twop(&allt, &ans, len(*wt), true, ans1)
	fmt.Fprintln(out, ans1)
	for _, val := range ans {
		fmt.Fprint(out, val, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for inp := 0; inp < t; inp++ {
		var n int
		fmt.Fscan(in, &n)
		th := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &th[i])
		}

		var m int
		fmt.Fscan(in, &m)
		wt := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &wt[i])
		}

		solve(out, &th, &wt)
	}
}
