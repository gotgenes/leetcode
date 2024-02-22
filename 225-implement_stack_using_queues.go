type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Size int
	head *Node
	tail *Node
}

func (q *Queue) Add(val int) {
	node := &Node{
		Value: val,
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
	val := q.head.Value
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

func (this *MyStack) Push(x int) {
    this.queue.Add(x)
}

func (this *MyStack) Pop() int {
    for i := 0; i < this.queue.Size-1; i++ {
        this.queue.Add(this.queue.Get())
    }
    return this.queue.Get()
}

func (this *MyStack) Top() int {
    for i := 0; i < this.queue.Size-1; i++ {
        this.queue.Add(this.queue.Get())
    }
    val := this.queue.Get()
    this.queue.Add(val)
    return val
}

func (this *MyStack) Empty() bool {
	return this.queue.Size == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */