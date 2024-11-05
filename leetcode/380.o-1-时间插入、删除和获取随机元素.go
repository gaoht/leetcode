package leetcode

import "math/rand"
type RandomizedSet struct {
    members map[int]int
    indices []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        members: map[int]int{},
        indices: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.members[val]; ok {
        return false
    }
    this.indices = append(this.indices, val)
    this.members[val] = len(this.indices) -  1
    return true

}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.members[val]; !ok {
        return false
    }
    id := this.members[val]
    lastVal := this.indices[len(this.indices) - 1]
    this.indices[id] = lastVal
    delete(this.members, val)
    this.indices = this.indices[ : len(this.indices) - 1]
    return false
}


func (this *RandomizedSet) GetRandom() int {
    return this.indices[rand.Intn(len(this.indices))]
}


