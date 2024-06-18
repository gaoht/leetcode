package array

func GenerateMatrix(n int) [][]int {
	a := make([][]int, n)
	length := n * n
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	top, bottom := 0, n-1
	left, right := 0, n-1
	i := 1
	for i <= length {
		for j := left; j <= right; j++ {
			a[top][j] = i
			i++
		}
		top ++
		
		for j := top; j <= bottom; j++ {
			a[j][right] = i
			i++
		}
		right --

		for j := right; j >= left; j--{
			a[bottom][j] = i
			i++	
		}
		bottom--

		for j:= bottom; j >= top; j-- {
			a[j][left] = i
			i++
		}
		left++
	}
	return a
}


func SpiralOrder(matrix [][]int) []int {
    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, n - 1
    top, bottom := 0, m - 1
    a := make([]int, m * n)
    i := 0
    for i < m*n {
		for j := left; j <= right; j++ {
			a[i] = matrix[top][j]
			i++
		}
		if (i == m*n){
			 break
		}
        top ++
    
		
        for j := top; j <= bottom; j++ {
            a[i] = matrix[j][right]
            i++
		}
		if (i == m*n){
			 break
		}
        right --

		if left <= right {
			for j := right; j >= left; j-- {
				a[i] = matrix[bottom][j]
				i++
			}
		}
		if (i == m*n){
			 break
		}
        bottom --

		if top <= bottom {
			for j := bottom; j >= top; j-- {
				a[i] = matrix[j][left]
				i++
			}
		}
		if (i == m*n){
			 break
		}
        left ++
    }
    return a
}