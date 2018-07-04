package main

import "fmt"

func main() {
	var s = []int{432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	shellSort(s)
	fmt.Println(s)
}

func shellSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j > 0 && array[j] < array[j-1] {
				exch(array, j, j-1)
				j = j - key
			}
		}
		key = key / 2
	}
}

func exch(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
