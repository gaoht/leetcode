package hash
// Intersection 349
func Intersection(nums1 []int, nums2 []int) []int {
    dict := [1001]int{}
    m := make(map[int]bool)
    for _, i := range nums1 {
        dict[i] = 1
    }
    for _, i := range nums2 {
        if dict[i] == 1 {
            m[i] = true
        }
    }
    a := make([]int, 0)
    for k := range m {
        a = append(a, k)
    }
    return a
}