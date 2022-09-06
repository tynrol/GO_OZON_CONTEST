package main

import (
	"bufio"
	"fmt"
	"os"
)

// map[value]count
var stack map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, a, res int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		res = 0
		stack = make(map[int]int)
		fmt.Fscan(in, &m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			stack[a] += 1
		}
		for value, count := range stack {
			res += value * (count - count/3)
		}
		fmt.Println(res)
	}
}
