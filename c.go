package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	coding := map[string]string{
		"00":  "a",
		"100": "b",
		"101": "c",
		"11":  "d",
	}

	in := bufio.NewReader(os.Stdin)
	var n int
	var str, res string
	var key string = ""
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str)
		for _, char := range str {
			if len(key) == 0 {
				key = string(char)
				continue
			}
			key += string(char)
			if _, found := coding[key]; found {
				res += coding[key]
				key = key[:0]
			}
		}
		fmt.Println(res)
		res = res[:0]
	}
}
