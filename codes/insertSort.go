package main

import (
	"fmt"
)

func main() {

	var s = []int{432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}

	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			swap(s, j, j-1) //这里是一个slice的地址拷贝
		}
	}

	fmt.Println(s)
}
func swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
