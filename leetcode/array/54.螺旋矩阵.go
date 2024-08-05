package leetcode
// [54] 螺旋矩阵

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	loop := len(matrix) * len(matrix[0])
	nums := make([]int, loop)
	left, right, top, bottom := 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	index := 0
	for loop > 0 {
		for i, j := top, left; j <= right; j, index = j + 1, index + 1{
			nums[index] = matrix[i][j] 
			loop --
		}
		top++
		if loop == 0 {
			break;
		}
		for i, j := top, right; i <= bottom; i, index = i + 1, index + 1 {
			nums[index] = matrix[i][j] 
			loop --
		}
		right--
		if loop == 0 {
			break;
		}
		for i, j := bottom, right; j >= left; j, index = j - 1, index + 1{
			nums[index] = matrix[i][j]
			loop --
		}
		if loop == 0 {
			break;
		}
		bottom--
		for i, j := bottom, left; i >= top; i, index = i - 1, index + 1{
			nums[index] = matrix[i][j]
			loop --
		}
		if loop == 0 {
			break;
		}
		left ++
	}
	return nums
}

