package problems

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)

	if len(ms.minStack) == 0 || ms.minStack[len(ms.minStack)-1] >= val {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) Pop() {
	val := ms.stack[len(ms.stack)-1]

	ms.stack = ms.stack[:len(ms.stack)-1]

	if ms.minStack[len(ms.minStack)-1] == val {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
