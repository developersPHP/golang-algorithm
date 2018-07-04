package main

import "fmt"

func main() {
	key := 12
	r := []int{1, 2, 3, 10, 12, 14}
	res := rank(key, r)
	fmt.Printf("最终输出结果为%d", res)

}

func rank(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	mid := lo + (hi-lo)/2
	fmt.Println(mid) //2

	for {
		if key < a[mid] {
			mid = mid - 1
		} else if key > a[mid] {
			mid = mid + 1
			fmt.Println(hi)
		} else {
			fmt.Printf("输出的最后结果为%d\n", mid)
			break
		}
	}
	return mid

}
