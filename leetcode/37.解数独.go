package leetcode

import "fmt"

var loops = 0

func solveSudoku(board [][]byte)  {
    var dfs func(cur int) bool
    dfs = func(cur int) bool{
        for i := cur; i < len(board) * len(board); i++ {
            loops ++
            row := i / len(board) 
            col := i % len(board)
            if board[row][col] == '.' {
                for c := byte('1'); c <= '9'; c++ {
                    if isValidSudoku(board, c, row, col) {
                        board[row][col] = c
                        if dfs(0) {
                            return true
                        }
                        board[row][col] = '.'
                    }
                }
                return false
            }
        }
        return true
    }
	dfs(0)
}

func solveSudoku2(board [][]byte)  {
    var dfs func(cur int) bool
    dfs = func(cur int) bool{
        for i := cur; i < len(board) * len(board); i++ {
            loops ++
            row := i / len(board)
            col := i % len(board)
            if board[row][col] != '.' {
                continue
            }
            for c := byte('1'); c <= '9'; c++ {
                if isValidSudoku(board, c, row, col) {
                    board[row][col] = c
                    if dfs(cur + 1) {
                        return true
                    }
                    board[row][col] = '.'
                } else {
                    continue
                }
            }
            return false
        }
        return true
    }
    dfs(0)
}

func isValidSudoku(board [][]byte, value byte, row, col int) bool {
    //同行
    for i := 0; i < col; i++ {
        if board[row][i] == value {
            return false
        }
    }
    //同列
    for i := 0; i < row; i++ {
        if board[i][col] == value {
            return false
        }
    }
    //九宫格
    startRow := row / 3 * 3
    startCol := col / 3 * 3
    for i := startRow; i < startRow + 3; i++ {
        for n := startCol; n < startCol + 3; n++{
            if (i != row || n != col) && board[i][n] == value {
                return false
            }
        }
    }
    return true
}

func printBoard(b  [][]byte){
	for i := 0; i < len(b); i ++ {
		for n := 0; n < len(b[i]); n++{
			fmt.Printf("%s ", string(b[i][n]))
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}