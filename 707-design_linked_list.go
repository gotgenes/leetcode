package main

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	return this.getNode(index).Value
}

func (this *MyLinkedList) getNode(index int) *Node {
	node := this.head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Value: val,
		Next:  this.head,
	}
	if this.head == nil {
		this.tail = node
	} else {
		this.head.Prev = node
	}
	this.head = node
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Value: val,
		Prev:  this.tail,
	}
	if this.tail == nil {
		this.head = node
	} else {
		this.tail.Next = node
	}
	this.tail = node
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else if index < this.length {
		node := this.getNode(index)
		newNode := &Node{
			Value: val,
			Prev:  node.Prev,
			Next:  node,
		}
		node.Prev.Next = newNode
		node.Prev = newNode
		this.length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < this.length {
		node := this.getNode(index)
		if index == 0 {
			this.head = node.Next
		}
		if index == this.length-1 {
			this.tail = node.Prev
		}
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		this.length--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

