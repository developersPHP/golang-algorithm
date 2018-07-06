package main

type Object interface {
}
type Node struct {
	data Object
	next *Node
}

type List struct {
	size int
	head *Node //链头
	tail *Node //链尾
}

func init() {
	(*List).size = 0
	(*List).head = nil
	(*List).tail = nil
}

func (this *List) Insert(i int, node *Node) bool {
	if node == nil || i > (*this).size || (*this).size == 0 {
		return false
	}

	if i == 0 {
		//直接排在首位
		(*node).next = (*this).head
		(*this).head = node
	} else {
		//排在指定的i位置的前面,找到排在i位置之前的node
		preItem := (*this).head
		for j := 1; j < i; j++ {
			preItem = (*preItem).next
		}
		//交换位置
		(*node).next = (*preItem).next
		(*preItem).next = preItem
	}
	(*this).size++
	return true
}

func (this *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	//给开辟空间
	(*node).next = nil
	if this.size == 0 {
		(*this).head = node
	} else {
		oldtail := (*this).tail
		oldtail.next = node
	}
	(*this).tail = node
	(*this).size++
	return true
}

func (this *List) Remove(i int, node *Node) bool {
	if i >= (*this).size {
		return false
	}
	if i == 0 {
		node = (*this).head
		(*this).head = (*node).next
		if (*this).size == 1 {
			(*this).tail = nil
		}
	} else {
		preItem := (*this).head
		for j := 1; j < i; j++ {
			preItem = (*preItem).next
		}
		node = (*preItem).next
		(*preItem).next = (*node).next
	}
	(*this).size--
	return true
}

func (this *List) Get(i int) *Node {
	if i >= (*this).size {
		return nil
	}
	item := (*this).head
	for j := 1; j < i; j++ {
		item = (*item).next
	}
	return item
}

func main() {

}
