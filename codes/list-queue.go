package main

import "fmt"

type Elem1 interface{}

type Node1 struct {
	data Elem1

	next *Node1
}

type QueueLink struct {
	front *Node1 // 对头

	tail *Node1 // 队尾

	length int
}

func (q *QueueLink) InitQueue() {

	q.front = new(Node1)

	q.tail = q.front

	q.length = 0

}

func (q *QueueLink) EnQueue(e Elem1) {

	Node1 := new(Node1)

	Node1.data = e

	Node1.next = nil

	q.tail.next = Node1

	q.tail = Node1

	q.length++

}

func (q *QueueLink) OutQueue() Elem1 {

	if q.Empty() {
		fmt.Println("queue is empty!")
	}

	N := q.front.next

	e := N.data

	q.front.next = N.next

	// 如果弹出的是队尾元素那么队列就空了，这个时候队尾等于队列

	if q.tail == N {

		q.tail = q.front

	}

	q.length--

	return e

}

func (q *QueueLink) Empty() bool {

	return q.front == q.tail

}

func (q *QueueLink) Length() int {

	return q.length

}

func (q *QueueLink) Destroy() {

	q.front = nil
	q.tail = nil
}

func main() {
	queue := new(QueueLink)
	queue.InitQueue()
	for i := 0; i < 50; i++ {
		queue.EnQueue(i)
	}

	fmt.Println("size:", queue.Length())
	fmt.Println("出队", queue.OutQueue())
	fmt.Println("size:", queue.Length())
}
