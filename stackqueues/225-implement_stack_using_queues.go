package stackqueues

import "github.com/gotgenes/leetcode"

type Queue struct {
	head *leetcode.ListNode
	tail *leetcode.ListNode
	Size int
}

func (q *Queue) Add(val int) {
	node := &leetcode.ListNode{
		Val: val,
	}
	if q.head == nil {
		q.head = node
	}
	if q.tail != nil {
		q.tail.Next = node
	}
	q.tail = node
	q.Size++
}

func (q *Queue) Get() int {
	val := q.head.Val
	q.head = q.head.Next
	q.Size--
	return val
}

type MyStack struct {
	queue *Queue
}

func Constructor() MyStack {
	return MyStack{queue: &Queue{}}
}

func (m *MyStack) Push(x int) {
	m.queue.Add(x)
}

func (m *MyStack) Pop() int {
	for i := 0; i < m.queue.Size-1; i++ {
		m.queue.Add(m.queue.Get())
	}
	return m.queue.Get()
}

func (m *MyStack) Top() int {
	for i := 0; i < m.queue.Size-1; i++ {
		m.queue.Add(m.queue.Get())
	}
	val := m.queue.Get()
	m.queue.Add(val)
	return val
}

func (m *MyStack) Empty() bool {
	return m.queue.Size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

