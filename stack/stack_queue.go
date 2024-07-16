package stack
// MyQueue 23
type MyQueue struct {
	m IntStack
	n IntStack
}


func Constructor() MyQueue {
	return MyQueue{
		m: *NewStack(),
		n: *NewStack(),
	}
}


func (q *MyQueue) Push(x int)  {
	q.m.Push(x)
}


func (q *MyQueue) Pop() int {
	if !q.n.IsEmpty() {
		return q.n.Pop()
	}
	for !q.m.IsEmpty() {
		q.n.Push(q.m.Pop())
	}
	return q.n.Pop();
}


func (q *MyQueue) Peek() int {
	e := q.Pop()
	q.n.Push(e)
	return e
}


func (q *MyQueue) Empty() bool {
	return q.m.IsEmpty() && q.n.IsEmpty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */