package leetcode
type Node struct{
    Pre *Node
    Next *Node
    Key int
    Val int
}
type LRUCache struct {
    head *Node
    dict map[int]*Node
    capacity int
}


func Constructor(capacity int) LRUCache {
    if capacity < 1 {
        capacity = 1
    }
    return LRUCache{
        dict: make(map[int]*Node, capacity),
        capacity: capacity,
    }   
}


func (this *LRUCache) Get(key int) int {
    node := this.dict[key]
    if node == nil {
        return -1
    }
    this.moveToHead(node)
    return node.Val
}


func (this *LRUCache) Put(key int, value int)  {
    node := this.dict[key]
    full := false
    if node == nil {
        node = &Node{
            Val: value,
            Key: key,
        }
        // 如果超过capacity了 需要移除一个
        full = this.capacity == len(this.dict) 
    } else {
        node.Val = value
    }
    if full {
        tail := this.head.Pre 
        delete(this.dict, tail.Key)
        if tail == this.head {
            this.head = nil
        } else {
            tail.Pre.Next = this.head
            this.head.Pre = tail.Pre
        }
    }
    this.moveToHead(node)
    this.dict[key] = node
    
}
func (this *LRUCache) moveToHead(node *Node){
    if node == this.head {
        return
    }
    if node.Pre != nil {
        node.Pre.Next = node.Next
        node.Next.Pre = node.Pre
    }
    h := this.head 
    if h == nil {
        this.head = node
        this.head.Pre = node
        this.head.Next = node
    } else {
        this.head = node
        // 处理新节点
        node.Pre = h.Pre
        node.Next = h

        // 处理尾节点
        h.Pre.Next = node
        // 处理原头结点
        h.Pre = node

    }
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
    */