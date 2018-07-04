### 算法一 选择性排序

```go
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

```

#### 特点
* 运行时间和输入无关，为了找出最小的元素而遍历一遍数组，但是上一次的变量不能为下一次遍历提供任何信息，造成一个现象就是一个已经排序好的数组和一个元素随机的数组
排序所用时间是一样的

* 数据移动最少。每次交换都是改变两个元素，交换的次数和数组长度成线性关系
