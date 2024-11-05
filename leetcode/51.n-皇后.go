package leetcode
func solveNQueens(n int) [][]string {
    ans := [][]string{}
    var dfs func(board []string)
    dfs = func(board []string) {
        if len(board) == n {
            tmp := make([]string, n)
            copy(tmp, board)
            ans = append(ans, tmp)
            return
        }
        for i := 0; i < n; i++ {
            row := buildRow(i, n)
            board = append(board, row)
            if isValidQueens(board, i, n) {
                dfs(board)
                if len(board) == n {
                    return
                }
            } 
            board = board[0 : len(board) - 1]
        }
    }
    dfs(make([]string, 0, n))
    return ans
}

func buildRow(col, n int) string {
    s := make([]byte, 0, n)
    for i := 0; i < n; i++{
        if i == col {
            s = append(s, 'Q')
        } else {
            s = append(s, '.')
        }
    }
    return string(s)
}

func isValidQueens(board []string, col, n int) bool{
    row := len(board) - 1 
    // 同列
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }
    // 45度
    for i := 0; i < row; i++ {
        c := col - (row - i)
        if col >= 0 {
            if board[i][c] == 'Q' {
                return false
            }
        }
    }
    // 135度
    for i := 0; i < row; i++ {
        c := col + (row - i)
        if col < n {
            if board[i][c] == 'Q' {
                return false
            }
        }
    }
    return true
}

