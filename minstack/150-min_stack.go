package minstack

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(val int) {
	m.values = append(m.values, val)
	if len(m.mins) == 0 {
		m.mins = append(m.mins, val)
	} else if lastMin := m.mins[len(m.mins)-1]; val < lastMin {
		m.mins = append(m.mins, val)
	} else {
		m.mins = append(m.mins, lastMin)
	}
}

func (m *MinStack) Pop() {
	m.values = m.values[:len(m.values)-1]
	m.mins = m.mins[:len(m.mins)-1]
}

func (m *MinStack) Top() int {
	return m.values[len(m.values)-1]
}

func (m *MinStack) GetMin() int {
	return m.mins[len(m.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
