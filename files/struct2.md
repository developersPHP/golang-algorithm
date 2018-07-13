### 数据结构二 栈
```go
package main

import (
	"fmt"
	"time"
)

//双向链表实现栈
type Elem interface{}

var header *entry //链表表头
var size int      //栈长度

type entry struct {
	previous *entry
	next     *entry
	element  Elem
}

func newEntry(prev, next *entry, e Elem) *entry {
	return &entry{prev, next, e}
}

//初始化header 表头
func NewStack() *entry {
	header = newEntry(nil, nil, nil)
	header.previous = header
	header.next = header
	return header
}

type Stack interface {
	Push(e Elem)   //向栈顶添加元素
	Pop() Elem     //移除栈顶弹出元素
	Top() Elem     //获取栈顶元素
	Clear() bool   //清空栈
	Size() int     //获取栈的元素个数
	IsEmpty() bool //判断栈是否是空栈
}

func (this *entry) Push(e Elem) {
	//创建一个栈元素
	newEntry := newEntry(header.previous, header, e)
	newEntry.previous.next = newEntry
	newEntry.next.previous = newEntry
	size++
	//fmt.Println(*newEntry)
	//fmt.Println(header)
	//fmt.Println("--------")

}

func (this *entry) Pop() Elem {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
		return nil
	}
	prevEntry := header.previous

	prevEntry.previous.next = header
	header.previous = prevEntry.previous

	size--
	return prevEntry.element
}

//获取栈顶元素(不删除)
func (this *entry) Top() Elem {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
		return nil
	}

	return this.previous.element
}

func (this *entry) Clear() bool {
	if this.IsEmpty() {
		fmt.Println("stack is empty")
		return false
	}
	entry := header.next
	//从下往上清除
	for entry != header {
		nextEntry := entry.next
		entry.next = nil
		entry.previous = nil
		entry.element = nil
		entry = nextEntry
	}

	header.next = header
	header.previous = header
	size = 0

	return true
}

func (this *entry) Size() int {
	return size
}

func (this *entry) IsEmpty() bool {
	if size == 0 {
		return true
	}
	return false
}

//****************************************
//****************************************
//用切片实现Stack
type sliceEntry struct {
	element []Elem
}

func NewSliceEntry() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Push(e Elem) {
	entry.element = append(entry.element, e)
}

func (entry *sliceEntry) Pop() Elem {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return nil
	}
	lastElement := entry.element[size-1]
	entry.element[size-1] = nil
	entry.element = entry.element[:size-1]
	return lastElement
}

func (entry *sliceEntry) Top() Elem {
	size := entry.Size()
	if size == 0 {
		fmt.Println("stack is empty!")
		return nil
	}
	return entry.element[size-1]
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("stack is empty!")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = make([]Elem, 0)
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func main() {
	test3()
}

func test1() {
	stack := NewStack()
	for i := 0; i < 50; i++ {
		stack.Push(i)
	}
	fmt.Println(stack.previous)
	fmt.Println(stack.Top())
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
	//fmt.Println(stack.Top())
	//fmt.Println(stack.Clear())
}

func test2() {
	entry := NewSliceEntry()
	for i := 0; i < 50; i++ {
		entry.Push(i)
	}
	fmt.Println(entry.Size())
	fmt.Println(entry.Top())
	fmt.Println(entry.Pop())
	fmt.Println(entry.Top(), entry.Size())
	fmt.Println(entry.Clear())
	for i := 0; i < 50; i++ {
		entry.Push(i)
	}
	fmt.Println(entry.Size())

}

//两种方法性能比较
func test3() {
	t := time.Now()
	sliceStack := NewSliceEntry()
	for i := 0; i < 500000; i++ {
		sliceStack.Push(i)
	}
	fmt.Println(time.Since(t)) //198.1407ms

	t = time.Now()
	stack := NewStack()
	for i := 0; i < 500000; i++ {
		stack.Push(i)
	}
	fmt.Println(time.Since(t)) //54.0385ms
}

```

#### 定义
>栈（stack）是限制插入和删除只能在一个位置上进行的线性表，该位置在表的末端，叫做栈顶。
添加元素只能在尾节点后添加，删除元素只能删除尾节点，查看节点也只能查看尾节点。
添加、删除、查看依次为入栈（push）、出栈（pop）、栈顶节点（top）。
形象的说，栈是一个先进后出（LIFO）表，先进去的节点要等到后边进去的节点出来才能出来。
如下图

![](https://github.com/developersPHP/golang-algorithm/blob/master/images/stack_1.png)

>如上图，是一个栈的形象图，top指针指向的是栈顶节点，所以我们可以通过top访问到2节点，但是0和1节点由于先于2进入这个表，所以是不可见的。如果把0节点当做头节点，2节点当做尾节点，那么栈限制了访问权限，只可以访问尾节点

#### 使用场景
>


