### 算法四 归并排序
```go
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

```

#### 定义
>将一个序列划分为同样大小的两个子序列,然后对两个子序列分别进行排序,最后进行合并操作,将两个子序列合成有序的序列.在合成的过程中,一般的实现都需要开辟一块与原序列大小相同的空间,以进行合并操作
>但是有一种空间复杂度为O(1)的原地归并算法，节省空间，原地归并可以将数组a分成两个数组a[begin...mid-1]和a[mid...end]并且归并成一个有序的数组，但是要注意，数组a[begin...mid-1]和a[mid...end]分别是有序的，这一点非常重要

#### 实现过程   
>归并排序有两部分组成，首先把无序的数据分成两个数组，分别通过>递归调用成两个有序的数组，然后通过原地归并算法进行归并操作，最终形成一个有序数组

>递归的array [432 432432 4234 333 333 21 22 3 30 8 20 2 7 9 50 80 1 4]

>递归计算key 9

>递归的array [432 432432 4234 333 333 21 22 3 30]

>递归计算key 4

>递归的array [432 432432 4234 333]

>递归计算key 2

>递归的array [432 432432]

>递归计算key 1

>递归的array [432]

>递归计算完left [432]

>此时的array [432 432432]

>递归的array [432432]

>递归计算完right [432432]

>递归的left [432]

>递归的right [432432]

>tmp最终结果为 [432 432432]

>递归计算完left [432 432432]

>此时的array [432 432432 4234 333]

>递归的array [4234 333]

>递归计算key 1

>递归的array [4234]

>递归计算完left [4234]

>此时的array [4234 333]

>递归的array [333]

>递归计算完right [333]

>递归的left [4234]

>递归的right [333]

>tmp最终结果为 [333 4234]

>递归计算完right [333 4234]

>递归的left [432 432432]

>递归的right [333 4234]

>tmp最终结果为 [333 432 4234 432432]

>递归计算完left [333 432 4234 432432]

>此时的array [432 432432 4234 333 333 21 22 3 30]

>递归的array [333 21 22 3 30]

>递归计算key 2

>递归的array [333 21]

>递归计算key 1

>递归的array [333]

>递归计算完left [333]

>此时的array [333 21]

>递归的array [21]

>递归计算完right [21]

>递归的left [333]

>递归的right [21]

>tmp最终结果为 [21 333]

>递归计算完left [21 333]

>此时的array [333 21 22 3 30]

>递归的array [22 3 30]

>递归计算key 1

>递归的array [22]

>递归计算完left [22]

>此时的array [22 3 30]

>递归的array [3 30]

>递归计算key 1

>递归的array [3]

>递归计算完left [3]

>此时的array [3 30]

>递归的array [30]

>递归计算完right [30]

>递归的left [3]

>递归的right [30]

>tmp最终结果为 [3 30]

>递归计算完right [3 30]

>递归的left [22]

>递归的right [3 30]

>tmp最终结果为 [3 22 30]

>递归计算完right [3 22 30]

>递归的left [21 333]

>递归的right [3 22 30]

>tmp最终结果为 [3 21 22 30 333]

>递归计算完right [3 21 22 30 333]

>递归的left [333 432 4234 432432]

>递归的right [3 21 22 30 333]

>tmp最终结果为 [3 21 22 30 333 333 432 4234 432432]

>递归计算完left [3 21 22 30 333 333 432 4234 432432]

>此时的array [432 432432 4234 333 333 21 22 3 30 8 20 2 7 9 50 80 1 4]

>递归的array [8 20 2 7 9 50 80 1 4]

>递归计算key 4

>递归的array [8 20 2 7]

>递归计算key 2

>递归的array [8 20]

>递归计算key 1

>递归的array [8]

>递归计算完left [8]

>此时的array [8 20]

>递归的array [20]

>递归计算完right [20]

>递归的left [8]

>递归的right [20]

>tmp最终结果为 [8 20]

>递归计算完left [8 20]

>此时的array [8 20 2 7]

>递归的array [2 7]

>递归计算key 1

>递归的array [2]

>递归计算完left [2]

>此时的array [2 7]

>递归的array [7]

>递归计算完right [7]

>递归的left [2]

>递归的right [7]

>tmp最终结果为 [2 7]

>递归计算完right [2 7]

>递归的left [8 20]

>递归的right [2 7]

>tmp最终结果为 [2 7 8 20]

>递归计算完left [2 7 8 20]

>此时的array [8 20 2 7 9 50 80 1 4]

>递归的array [9 50 80 1 4]

>递归计算key 2

>递归的array [9 50]

>递归计算key 1

>递归的array [9]

>递归计算完left [9]

>此时的array [9 50]

>递归的array [50]

>递归计算完right [50]

>递归的left [9]

>递归的right [50]

>tmp最终结果为 [9 50]

>递归计算完left [9 50]

>此时的array [9 50 80 1 4]

>递归的array [80 1 4]

>递归计算key 1

>递归的array [80]

>递归计算完left [80]

>此时的array [80 1 4]

>递归的array [1 4]

>递归计算key 1

>递归的array [1]

>递归计算完left [1]

>此时的array [1 4]

>递归的array [4]

>递归计算完right [4]

>递归的left [1]

>递归的right [4]

>tmp最终结果为 [1 4]

>递归计算完right [1 4]

>递归的left [80]

>递归的right [1 4]

>tmp最终结果为 [1 4 80]

>递归计算完right [1 4 80]

>递归的left [9 50]

>递归的right [1 4 80]

>tmp最终结果为 [1 4 9 50 80]

>递归计算完right [1 4 9 50 80]

>递归的left [2 7 8 20]

>递归的right [1 4 9 50 80]

>tmp最终结果为 [1 2 4 7 8 9 20 50 80]

>递归计算完right [1 2 4 7 8 9 20 50 80]

>递归的left [3 21 22 30 333 333 432 4234 432432]

>递归的right [1 2 4 7 8 9 20 50 80]

>tmp最终结果为 [1 2 3 4 7 8 9 20 21 22 30 50 80 333 333 432 4234 432432]
[1 2 3 4 7 8 9 20 21 22 30 50 80 333 333 432 4234 432432]

#### 应用场景

#### 特点




