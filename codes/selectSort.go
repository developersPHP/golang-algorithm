package main

import "fmt"

func main() {
	var e = []int{2, 3, 1, 4, 6}
	res := sort(e)
	fmt.Printf("排序后的结果为%d", res)
}

func sort(s []int) []int {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[j] < s[i] {
				tmp := s[i]
				s[i] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}
