package stack

type Queue struct{
	queue []int
	size int
}

func (q *Queue) Push(n int){
	q.queue = append(q.queue, n)
	q.size++
}

func (q *Queue) Peek(n int) int{
	return q.queue[0]
}


func (q *Queue) Pop(n int) int{
	e := q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return e
}

func (q *Queue) Size() int{
	return q.size
}

func (q *Queue) IsEmpty() bool{
	return q.size == 0
}