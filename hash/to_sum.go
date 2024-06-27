package hash
// TwoSum 1
func TwoSum2(nums []int, target int) []int {
    for i := 0; i < len(nums); i++{
        for j := i+1; j < len(nums); j++{
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}
func TwoSum(nums []int, target int) []int {
    a := make(map[int]int)
	for k, v := range nums {
		a[target - v] = k
	}
	for k, v := range nums{
		if i, ok := a[v]; ok && i != k  {
			return []int{i,k}
		}
	} 
	return []int{}
}