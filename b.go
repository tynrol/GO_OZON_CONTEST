package main

import (
	"bufio"
	"fmt"
	"os"
)

//i dont like this very much
func main() {
	in := bufio.NewReader(os.Stdin)
	var n, day, month, year int
	var visok bool = false
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &day, &month, &year)
		visok = false
		if (year%400 == 0) || (year%4 == 0 && !(year%100 == 0)) {
			visok = true
		}
		if month == 2 {
			if visok {
				if day > 29 {
					fmt.Println("NO")
					continue
				}
			} else {
				if day > 28 {
					fmt.Println("NO")
					continue
				}
			}
			fmt.Println("YES")
			continue
		}
		if (month <= 7 && month%2 == 1) || (month > 7 && month%2 == 0) {
			if day > 31 {
				fmt.Println("A")
				continue
			}
		} else {
			if day > 30 {
				fmt.Println("NO")
				continue
			}
		}
		fmt.Println("YES")
	}
}
