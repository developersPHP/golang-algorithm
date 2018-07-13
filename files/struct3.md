### 数据结构 队列

#### 定义
>队列同样是一种特殊的线性表，其插入和删除的操作分别在表的两端进行，队列的特点就是先进先出(First In First Out)。我们把向队列中插入元素的过程称为入队(Enqueue)，删除元素的过程称为出队(Dequeue)并把允许入队的一端称为队尾，允许出的的一端称为队头，没有任何元素的队列则称为空队
>循环队列：采用环状顺序表来存放队列元素，并用两个指针，其中 front 指针指向队列的队头元素，rear指针指向队尾元素的下一个位置，往队列中加进或取出元素时分别改变这两个变量的计数。当头尾指针（front / rear）指向队列尾的元素（下标：QueueSize-1）时，其加1操作的结果是指向向量的下界0
>队列种类：单链队列、循环队列、阵列队列
![](https://github.com/developersPHP/golang-algorithm/blob/master/images/queue1.png)

#### 实现原理
* 顺序队列设计
> 关于顺序队列（底层都是利用数组作为容器）的实现，我们将采用顺序循环队列的结构来实现，在给出实现方案前先来分析一下为什么不直接使用顺序表作为底层容器来实现。实际上采用顺序表实现队列时，入队操作直接执行顺序表尾部插入操作，其时间复杂度为O(1)，出队操作直接执行顺序表头部删除操作，其时间复杂度为O(n)，主要用于移动元素，效率低，既然如此，我们就把出队的时间复杂度降为O(1)即可，为此在顺序表中添加一个头指向下标front和尾指向下标，出队和入队时只要改变front、rear的下标指向取值即可，此时无需移动元素，因此出队的时间复杂度也就变为O(1)。其过程如下图所示

![](https://github.com/developersPHP/golang-algorithm/blob/master/images/queue2.png)
```go
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

func NewQueue() *sliceEntry1  {
  return &sliceEntry1{}
}

//向队列中添加元素
func (entry *sliceEntry1) Offer(e Element1)  {
  entry.element = append(entry.element, e)
}

//移除队列中最前面的元素
func (entry *sliceEntry1) Poll() Element1  {
  if entry.IsEmpty() {
    fmt.Println("queue is empty!")
    return nil
  }
  
  firstElement := entry.element[0]
  entry.element = entry.element[1:]
  return firstElement
}

func (entry *sliceEntry1) Clear() bool  {
  if entry.IsEmpty() {
    fmt.Println("queue is empty!")
    return false
  }
  
  for i:=0; i < entry.Size(); i++ {
    entry.element[i] = nil
  }
  entry.element = nil
  return true
}

func (entry *sliceEntry1) Size() int  {
  return len(entry.element)
}

func (entry *sliceEntry1) IsEmpty() bool  {
  if len(entry.element) == 0 {
    return true
  }
  return false
}

func main()  {
  queue := NewQueue()
  for i:=0; i < 50; i++ {
    queue.Offer(i)
  }
  
  fmt.Println("size:", queue.Size())
  fmt.Println("移除最前面的元素:", queue.Poll())
  fmt.Println("size:",queue.Size())
  fmt.Println("清空",queue.Clear())
  
}


````

>从图的演示过程，（a）操作时，是空队列此时front和rear都为-1，同时可以发现虽然我们通过给顺序表添加front和rear变量记录下标后使用得出队操作的时间复杂度降为O(1)，但是却出现了另外一个严重的问题，那就是空间浪费，从图中的（d）和（e）操作可以发现，20和30出队后，遗留下来的空间并没有被重新利用，反而是空着，所以导致执行（f）操作时，出现队列已满的假现象，这种假现象我们称之为假溢出，之所以出现这样假溢出的现象是因为顺序表队列的存储单元没有重复利用机制，而解决该问题的最合适的方式就是将顺序队列设计为循环结构，接下来我们就通过循环顺序表来实现顺序队列。 
顺序循环队列就是将顺序队列设计为在逻辑结构上收尾相接的循环结构，这样我们就可以重复利用存储单元，其过程如下所示
 
![](https://github.com/developersPHP/golang-algorithm/blob/master/images/queue3.png)
 
* 链式队列设计
>对于链式队列，将使用带头指针front和尾指针rear的单链表实现，front直接指向队头的第一个元素，rear指向队尾的最后一个元素，其结构如下
![](https://github.com/developersPHP/golang-algorithm/blob/master/images/queue4.png)

>之所以选择单链表（带头尾指针）而不采用循环双链表或者双链表主要是双链表的空间开销（空间复杂度，多前继指针）相对单链表来说大了不少，而单链表只要新增头指针和尾指针就可以轻松实现常数时间内（时间复杂度为O(1)）访问头尾结点。下面我们来看看如何设计链式队列：
* 以上述的图为例分别设置front和rear指向队头结点和队尾结点，使用单链表的头尾访问时间复杂度为O(1)。
* 设置初始化空队列，使用front=rear=null，并且约定条件front==null&&rear==null成立时，队列为空。
* 出队操作时，若队列不为空获取队头结点元素，并删除队头结点元素，更新front指针的指向为front=front.next
* 入队操作时，使插入元素的结点在rear之后并更新rear指针指向新插入元素。
* 当第一个元素入队或者最后一个元素出队时，同时更新front指针和rear指针的指向
>实现过程如下图
![](https://github.com/developersPHP/golang-algorithm/blob/master/images/queue5.png)
```go
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

````

#### 应用
