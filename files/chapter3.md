### 算法三 希尔排序(递减增量排序算法)

```go
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

```

#### 定义
>使一个随机数组中，用任意间隔的h划分为多个h数组，并且让h数组有序，最后通过插入排序进行排序，提高插入排序效率。
希尔排序主要是针对插入排序中两个缺点进行改进，第一，插入排序每次只能移动一个元素，第二，插入排序在几乎已经排好序的数组中操作时，效率最高
希尔排序是插入排序的一种，其原理是基于插入排序，除了希尔排序，还有伸展排序、二叉查找树排序、图书馆排序、耐心排序

#### 实现过程
>例如，假设有这样一组数[ 13 14 94 33 82 25 59 94 65 23 45 27 73 25 39 10 ]，如果我们以步长为5开始进行排序，我们可以通过将这列表放在有5列的表中来更好地描述算法，这样他们就应该看起来是这样：
 
 >13 14 94 33 82
 25 59 94 65 23
 45 27 73 25 39
 10
 然后我们对每列进行排序：
 
 >10 14 73 25 23
 13 27 94 33 39
 25 59 94 65 82
 45
 将上述四行数字，依序接在一起时我们得到：[ 10 14 73 25 23 13 27 94 33 39 25 59 94 65 82 45 ].这时10已经移至正确位置了，然后再以3为步长进行排序：
 
 >10 14 73
 25 23 13
 27 94 33
 39 25 59
 94 65 82
 45
 排序之后变为：
 
 >10 14 13
 25 23 33
 27 25 59
 39 65 73
 45 94 82
 94
 最后以1步长进行排序（此时就是简单的插入排序了）。
 
 #### 应用场景
 >主要用于无序的数组排序，可用插入排序的场景都能用希尔排序。
 #### 特点
 * 步长的选择是希尔排序的重要部分。只要最终步长为1任何步长序列都可以工作。算法最开始以一定的步长进行排序。然后会继续以一定步长进行排序，最终算法以步长为1进行排序。当步长为1时，算法变为插入排序，这就保证了数据一定会被排序。
 * 希尔排序最重要的地方在于当用较小步长排序后，以前用的较大步长仍然是有序的
 * 已知的最好步长序列是由Sedgewick提出的(1, 5, 19, 41, 109,...)
 