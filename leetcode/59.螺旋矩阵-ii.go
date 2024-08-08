package leetcode
// [59] 螺旋矩阵 II


func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++{
		m[i] = make([]int, n)
	}
	loop := n * n
	cursor := 1
	left, right, top, bottom := 0, n - 1, 0, n -1
	for loop > 0 {
		for i, j := top, left; j <= right; j++{
			m[i][j] = cursor
			cursor++
			loop--
		}
		top++
		if loop == 0 {
			break
		}
		for i, j := top, right; i <= bottom; i++{
			m[i][j] = cursor
			cursor++
			loop--
		}
		right--
		if loop == 0 {
			break
		}
		for i, j := bottom, right; j >= left; j--{
			m[i][j] = cursor
			cursor++
			loop--
		}
		bottom--
		if loop == 0 {
			break
		}
		for i, j := bottom, left; i >= top; i--{
			m[i][j] = cursor
			cursor++
			loop--
		}
		left++
	}
	return m
}

