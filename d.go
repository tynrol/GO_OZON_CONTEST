package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, value int
	var a1, a2, max, res int
	var input []int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &value)
			input = append(input, value)
		}

		if len(input) < 3 {
			fmt.Println(len(input))
			input = input[:0]
			continue
		}

		max = 0

		for l := 0; l < m-1; l++ {
			res = 2

			a1 = input[l]
			a2 = input[l+1]

			for a1 == a2 && l < m-2 {
				res++
				l++
				a2 = input[l+1]
			}

			for k := l + 2; k < m; k++ {
				if (input[k] == a1) || (input[k] == a2) {
					res++
				} else {
					break
				}
			}

			if res > max {
				max = res
			}
		}
		input = input[:0]

		fmt.Println(max)
	}
}
