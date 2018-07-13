package main

import "fmt"

type Element1 interface {
}

type Queue interface {
	Offer(e Element1) //向队列中添加元素
	Poll() Element1   //移除队列中最前面的元素
	Clear() bool      //清空队列
	Size() int        //获取队列的元素个数
	IsEmpty() bool    //判断队列是否是空
}

type sliceEntry1 struct {
	element []Element1
}

func NewQueue() *sliceEntry1 {
	return &sliceEntry1{}
}

//向队列中添加元素
func (entry *sliceEntry1) Offer(e Element1) {
	entry.element = append(entry.element, e)
}

//移除队列中最前面的元素
func (entry *sliceEntry1) Poll() Element1 {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}

	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

func (entry *sliceEntry1) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return false
	}

	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry1) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry1) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

func main() {
	queue := NewQueue()
	for i := 0; i < 50; i++ {
		queue.Offer(i)
	}

	fmt.Println("size:", queue.Size())
	fmt.Println("移除最前面的元素:", queue.Poll())
	fmt.Println("size:", queue.Size())
	fmt.Println("清空", queue.Clear())

}
