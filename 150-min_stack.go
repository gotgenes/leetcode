package main

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else if lastMin := this.mins[len(this.mins)-1]; val < lastMin {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins, lastMin)
	}
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

