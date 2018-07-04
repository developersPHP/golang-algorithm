### 算法五 快速排序

```go
package main

import "fmt"

func main() {
	var e = []int{432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
	qsort(e)
	fmt.Printf("排序后的结果为%d", e)
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid, i := data[0], 1
	head, tail := 0, len(data)-1
	for head < tail {
		//fmt.Println(data, head, tail, i)
		if data[i] > mid {
			// 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			data[i], data[tail] = data[tail], data[i]
			tail-- // 尾下标左移一位
		} else {
			// 如果分水岭右侧的元素小于等于分水岭，就将分水岭右移一位
			data[i], data[head] = data[head], data[i]
			head++ // 头下标右移一位
			i++    // i下标右移一位
		}
		fmt.Printf(">{head:%d,tail:%d,i:%d,mid:%d} data:%v\n\n", head, tail, i, mid, data)
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

```

#### 定义
