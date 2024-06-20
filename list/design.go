package list

// 707 LinkedListNode
type LinkedListNode struct{
	Val int
	Next *LinkedListNode
}
type MyLinkedList struct {
	Head *LinkedListNode 
	Size int
}


func Constructor() MyLinkedList {
	return MyLinkedList{Size: 0}
}


func (t *MyLinkedList) Get(index int) int {
	if index >= t.Size {
		return -1
	}
	vHead := &LinkedListNode{Next: t.Head}
	cur := vHead
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}


func (t *MyLinkedList) AddAtHead(val int)  {
	head := &LinkedListNode{
		Val: val,
		Next: t.Head,
	}
	t.Head = head
	t.Size++
}


func (t *MyLinkedList) AddAtTail(val int)  {
	vHead := &LinkedListNode{Next: t.Head}
	cur := vHead
	for cur.Next != nil{
		cur = cur.Next 
	}
	cur.Next = &LinkedListNode{
		Val: val,
	}
	t.Size ++
    t.Head = vHead.Next 
}


func (t *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > t.Size {
		return
	}
	vHead := &LinkedListNode{Next: t.Head}
	pre := vHead
	cur := vHead
	for i := 0; i <= index; i++ {
		pre = cur
		cur = cur.Next
	}
	tmp := pre.Next
	pre.Next = &LinkedListNode{
		Next: tmp,
		Val: val,
	}
	t.Size++
	t.Head = vHead.Next
}


func (t *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= t.Size {
		return
	}
	if t.Size == 0 {
		return
	}
	vHead := &LinkedListNode{
		Next: t.Head,
	}
	cur := vHead
	pre := vHead
	for i := 0; i <= index; i++{
		pre = cur
		cur = cur.Next 
	}
	pre.Next = cur.Next
	t.Size --
	t.Head = vHead.Next
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


// todo 双向链表未通过

type DLinkedListNode struct{
	Val int
	Next *DLinkedListNode
	Pre *DLinkedListNode
}
type MyDLinkedList struct {
	Head *DLinkedListNode 
	Tail *DLinkedListNode 
	Size int
}


func DConstructor() MyDLinkedList {
	return MyDLinkedList{Size: 0}
}


func (t *MyDLinkedList) Get(index int) int {
	if index >= t.Size {
		return -1
	}
	vHead := &DLinkedListNode{Next: t.Head}
	cur := vHead
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}


func (t *MyDLinkedList) AddAtHead(val int)  {
	head := &DLinkedListNode{
		Val: val,
	}
	if t.Head == nil {
		t.Tail = head
	} else {
		head.Next = t.Head
		t.Head.Pre = head
	}
	t.Head = head
	t.Size++
}


func (t *MyDLinkedList) AddAtTail(val int)  {
	if t.Tail == nil {
		t.AddAtHead(val)
	}
	tail := &DLinkedListNode{
		Val: val,
		Pre: t.Tail,
	}
	t.Tail.Next = tail
	t.Tail = tail
	t.Size++
}


func (t *MyDLinkedList) AddAtIndex(index int, val int)  {
	if index > t.Size {
		return
	}
	vHead := &DLinkedListNode{Next: t.Head}
	cur := vHead
	t.Head.Pre = vHead
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	n := &DLinkedListNode{
		Val: val,
		Pre: cur.Pre,
		Next: cur,
	}
	cur.Pre.Next = n
	cur.Pre = n
	if index == t.Size {
		t.Tail = cur
	}
	t.Size++
	t.Head = vHead.Next

}


func (t *MyDLinkedList) DeleteAtIndex(index int)  {
	if index >= t.Size {
		return
	}
	if t.Size == 0 {
		return
	}
	vHead := &DLinkedListNode{
		Next: t.Head,
	}
	cur := vHead
	t.Head.Pre = vHead
	for i := 0; i <= index; i++{
		cur = cur.Next 
	}
	cur.Pre.Next = cur.Next
	if index == t.Size -1 {
		t.Tail = cur
	}
	t.Size --
	t.Head = vHead.Next
}