package stack

type SlideQueue struct{
	q []int
}
func NewSlideQueue(l int) *SlideQueue{
	return &SlideQueue{
		q: make([]int, l),
	}
}
func (s *SlideQueue) IsEmpty() bool{
	return len(s.q) == 0
}

func (s *SlideQueue) Pop (n int){
	if !s.IsEmpty() && s.Front() != n  {
		s.q = s.q[1:]
	}
}

func (s *SlideQueue) Push(n int) {
	for !s.IsEmpty() && s.Back() < n  {
		s.q = s.q[ : len(s.q) - 1]
	}
	s.q = append(s.q, n)
}

func (s *SlideQueue) Front() int{
	return s.q[0]
}

func (s *SlideQueue) Back() int{
	return s.q[len(s.q) - 1]
}

// MaxSlideWindow 239
func MaxSlideWindow(nums []int, k int)[] int{
	s := NewSlideQueue(len(nums))
	for i := 0; i < k; i++ {
		s.Push(nums[i])
	}
	res := make([]int, len(nums) - k + 1)
	res = append(res, s.Front())
	for i := k; i < len(nums); i++{
		s.Pop(nums[i-k])
		s.Push(nums[i])
		res = append(res, s.Front())
	}
	return res		
}