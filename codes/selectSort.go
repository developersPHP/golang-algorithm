package main

import (
	"fmt"
)

func main() {
	var e = []int{432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	res := sort(e)
	fmt.Printf("排序后的结果为%d", res)
}

func sort(s []int) []int {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[j] < s[i] {
				exchange(s, j, i)
			}
		}
	}
	return s
}

func exchange(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
