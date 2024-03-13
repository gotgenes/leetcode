package linkedlist

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
}

type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (m *MyLinkedList) Get(index int) int {
	if index >= m.length {
		return -1
	}
	return m.getNode(index).Value
}

func (m *MyLinkedList) getNode(index int) *Node {
	node := m.head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (m *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Value: val,
		Next:  m.head,
	}
	if m.head == nil {
		m.tail = node
	} else {
		m.head.Prev = node
	}
	m.head = node
	m.length++
}

func (m *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Value: val,
		Prev:  m.tail,
	}
	if m.tail == nil {
		m.head = node
	} else {
		m.tail.Next = node
	}
	m.tail = node
	m.length++
}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		m.AddAtHead(val)
	} else if index == m.length {
		m.AddAtTail(val)
	} else if index < m.length {
		node := m.getNode(index)
		newNode := &Node{
			Value: val,
			Prev:  node.Prev,
			Next:  node,
		}
		node.Prev.Next = newNode
		node.Prev = newNode
		m.length++
	}
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index >= 0 && index < m.length {
		node := m.getNode(index)
		if index == 0 {
			m.head = node.Next
		}
		if index == m.length-1 {
			m.tail = node.Prev
		}
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}
		m.length--
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
