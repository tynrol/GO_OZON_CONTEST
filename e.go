package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Value, Left, Right int
}

var points map[Point]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	var a, b, c int
	var arr []int

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		points = make(map[Point]bool, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a, &b, &c)
			points[Point{a, b, c}] = false
		}
		arr = make([]int, m)
		array := solve(arr, m, true)
		array2 := solve(array, m, false)
		printPairs(array2)
	}

}

func solve(arr []int, m int, flag bool) []int {

	for key, val := range points {
		if val {
			continue
		}
		if flag {
			arr[0] = key.Value
			arr[1] = key.Right
			arr[len(arr)-1] = key.Left
			flag = false
			points[key] = true
			continue
		}
		for i := 0; i < m-1; i++ {
			if arr[i] == key.Value && arr[prevIndex(m, i)] == key.Right {
				arr[i+1] = key.Left
				points[key] = true
				continue
			}
			if arr[i] == key.Value && arr[prevIndex(m, i)] == key.Left {
				arr[i+1] = key.Right
				points[key] = true
				continue
			}
			if arr[i] == key.Value && arr[i+1] == key.Right {
				arr[prevIndex(m, i)] = key.Left
				points[key] = true
				continue
			}
			if arr[i] == key.Value && arr[i+1] == key.Left {
				arr[prevIndex(m, i)] = key.Right
				points[key] = true
				continue
			}
		}
	}

	return arr
}

func printPairs(array []int) {
	for i := 0; i < len(array)/2; i++ {
		fmt.Printf("%d %d\n", array[i], array[len(array)/2+i])
	}
}

func prevIndex(len, value int) int {
	if value == 0 {
		return len - 1
	} else {
		return value - 1
	}
}
