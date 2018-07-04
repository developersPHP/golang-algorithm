package main

import "fmt"

func main() {
	var e = []int{432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	arr := MergeSort(e)
	fmt.Println(arr)

}
func MergeSort(array []int) []int {
	fmt.Println("递归的array", array)
	l := len(array)
	if l < 2 {
		return array
	}
	key := l / 2
	fmt.Println("递归计算key", key)
	left := MergeSort(array[0:key]) //最终获取到的left是array [432]
	fmt.Println("递归计算完left", left)
	fmt.Println("此时的array", array)
	right := MergeSort(array[key:]) //最终获取到的right是array [432432]
	fmt.Println("递归计算完right", right)
	return merge(left, right)
}

//原地归并算法
func merge(left, right []int) []int {
	fmt.Println("递归的left", left)
	fmt.Println("递归的right", right)
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	fmt.Println("tmp最终结果为", tmp)
	return tmp
}
