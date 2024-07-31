package tree




type MaxHeap struct{
	nums []int
	count int
}

func NewMaxHeap() *MaxHeap{
	return &MaxHeap{
		nums: make([]int, 1),
	}

}

func (m *MaxHeap) shiftUp(n int){
	for n > 1 && m.nums[parent(n)] < m.nums[n] {
		swap(m.nums, parent(n), n)
		n = parent(n)
	}
}

func (m *MaxHeap) shiftDown(n int) {
	for left(n) < len(m.nums){
		i := left(n)
		if right(n) < len(m.nums) && m.nums[right(n)] > m.nums[left(n)]{
			i = right(n)
		}
		if m.nums[n] >= m.nums[i] {
			break
		}
		swap(m.nums, n, i)
		n = i
	}
}

func (m *MaxHeap)  Insert(num int) {
	m.nums = append(m.nums, num)
	m.shiftUp(len(m.nums) - 1)
	m.count++
}

func (m *MaxHeap) ExtractMax() (int, bool)  {
	if m.count == 0 {
		return 0, false
	}
	max := m.nums[1]
	swap(m.nums, 1, len(m.nums) - 1)
	m.nums = m.nums[: len(m.nums) - 1]
	m.count--
	m.shiftDown(1)
	return max, true
}

func swap(nums []int, n1 int, n2 int){
	nums[n1], nums[n2] = nums[n2], nums[n1]
}

func parent(n int) int {
	return n / 2
}

func left(n int) int{
	return 2 * n
}

func right(n int) int {
	return 2 * n + 1
}


type MinHeap struct {
	nums []int
	count int
}
func NewMinHeap() *MinHeap{
	return &MinHeap{
		nums: make([]int, 1),
		count: 0,
	}
}
func (h *MinHeap) shiftUp(n int){
	for n > 1 && h.nums[parent(n)] > h.nums[n] {
		swap(h.nums, parent(n), n)
		n = parent(n)
	}
}

func (h *MinHeap) shiftDown(n int){
	for left(n) < len(h.nums) {
		i := left(n)
		if right(n) < len(h.nums) && h.nums[right(n)] < h.nums[left(n)]{
			i = right(n)
		}
		if h.nums[n] > h.nums[i] {
			swap(h.nums, i, n)
			n = i
		} else {
			break
		}
	}
}
func (h *MinHeap) ExtractMin() (int, bool)  {
	if len(h.nums) <= 1 {
		return 0, false
	}
	n := h.nums[1]
	h.nums[1] = h.nums[len(h.nums) - 1]
	h.nums = h.nums[: len(h.nums) - 1]
	h.shiftDown(1)
	h.count -- 
	return n, true
}
func (h *MinHeap) Insert(n int){
	h.nums = append(h.nums, n)
	h.count ++
	h.shiftUp(h.count)
}

