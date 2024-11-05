package leetcode

import (
	"fmt"
	"sort"
)

type peopleSlice [][]int
func (p peopleSlice) Less(i, j int) bool {
    if p[i][0] > p[j][0] {
        return true
    } else if p[i][0] == p[j][0] {
        return p[i][1] < p[j][1]
    }
    return false
}
func (p peopleSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p peopleSlice) Len() int {
    return len(p)
}
func reconstructQueue(people [][]int) [][]int {
    sort.Sort(peopleSlice(people))
    var ans [][]int
    for i := 0; i < len(people); i++ {
		fmt.Println(ans, people[i])
        index := people[i][1]
        ans0 := ans[0 : index]
        ans1 := make([][]int, len(ans[index : ]))
		copy(ans1, ans[index : ])
        ans = append(ans0, people[i])
        if len(ans1) > 0 {
            ans = append(ans, ans1...)
        }
    }
    return ans
}