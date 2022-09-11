package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var str string
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		_, err := time.Parse("02 01 2006", str)
		if err != nil {
			fmt.Println("NO")
		}
		fmt.Println()
	}
}
