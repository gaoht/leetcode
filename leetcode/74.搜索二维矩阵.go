package leetcode

import "fmt"

//[74] 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
    left := -1
    m := len(matrix)
    n := len(matrix[0])
    right := m * n
    for left + 1 != right {
        mid := left + (right - left) / 2
        row, column := getPosition(mid, m, n)
		// fmt.Printf("%d, %d, %d \n", mid, row, column)
        if matrix[row][column] < target {
            left = mid
        } else if matrix[row][column] > target {
            right = mid
        } else {
            return true
        }
    }
    return false
}

func getPosition(mn, m, n int) (int, int){
    row := mn / n
    column := mn % n
	fmt.Printf("1111 %d, %d, %d, %d, %d \n", mn, m, n, row, column)
    return row, column
}