package LeetCode

import "fmt"

func Code37() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}

/**
编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sudoku-solver
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

```
*/
func solveSudoku(board [][]byte) {
	dfs_37(board, 0)

	fmt.Println(board)
}

/* k 是把 board 转换成一维数组后，元素的索引值 */
func dfs_37(board [][]byte, k int) bool {
	if k == 81 {
		return true
	}

	r, c := k/9, k%9
	if board[r][c] != '.' {
		return dfs_37(board, k+1)
	}

	/* bi, bj 是 rc 所在块的左上角元素的索引值 */
	bi, bj := r/3*3, c/3*3

	// 按照数独的规则，检查 b 能否放在 board[r][c]
	isValid := func(b byte) bool {
		for n := 0; n < 9; n++ {
			if board[r][n] == b ||
				board[n][c] == b ||
				board[bi+n/3][bj+n%3] == b {
				return false
			}
		}
		return true
	}

	for b := byte('1'); b <= '9'; b++ {
		if isValid(b) {
			board[r][c] = b
			if dfs_37(board, k+1) {
				return true
			}
		}
	}

	board[r][c] = '.'

	return false
}

/*func solveSudoku2(board [][]byte) {

	boardMap := make([][]map[byte]int, len(board))
	for i := range boardMap {
		boardMap[i] = make([]map[byte]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		eachMap := make(map[byte]int, 0)
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				eachMap[board[i][j]] = 1
			} else {
				eachMap[board[i][j]] = 0
			}
			boardMap[i][] = append(boardMap[i], eachMap)
		}
	}

	fmt.Printf("%+v\n", boardMap)
	solve := false
	dfs_37(board, 0, 0, boardMap, solve)

	fmt.Println(board)
}

func dfs_37(board [][]byte, x int, y int, boardMap [][]map[byte]int, solve bool) {
	if solve {
		return
	}
	if x >= 9 {
		solve = true
		return
	}
	if board[x][y] != '.' {
		if y < 8 {
			dfs_37(board, x, y+1, boardMap, solve)
		} else if y == 8 {
			dfs_37(board, x+1, 0, boardMap, solve)
		}
		return
	} else {
		for key := 1; key <= 9; key++ {
			fmt.Println(key, key-'0')
			if boardMap[x][y][byte(key)] == 1 {
				continue
			}
			if check(board, x, y, 48+byte(key)) {
				board[x][y] = '0' + byte(key)
				if _, ok := boardMap[x][y][byte(key)]; ok {

				}
				boardMap[x][y][byte(key)] = 1

				fmt.Println(board)
				if y < 8 {
					dfs_37(board, x, y+1, boardMap, solve)
					return
				} else if y == 8 {
					dfs_37(board, x+1, 0, boardMap, solve)
					return
				}
				if !solve {
					board[x][y] = '.'
					boardMap[x][y][byte(key)] = 0
				}
			}
		}
	}
}

func check(board [][]byte, x int, y int, s byte) bool {
	for i := 1; i <= 9; i++ {
		if board[x][i-1] == s {
			return false
		}
		if board[i-1][y] == s {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			curX, curY := begin(x, y)
			if board[curX+i][curY+j] == s {
				return false
			}
		}
	}
	return true
}

func begin(i int, j int) (int, int) {
	if i < 3 && j < 3 {
		return 0, 0
	}
	if i < 3 && j >= 3 && j < 6 {
		return 0, 3
	}
	if i < 3 && j >= 6 && j < 9 {
		return 0, 6
	}
	if i >= 3 && i < 6 && j < 3 {
		return 3, 0
	}
	if i >= 3 && i < 6 && j >= 3 && j < 6 {
		return 3, 3
	}
	if i >= 3 && i < 6 && j >= 6 && j < 9 {
		return 3, 6
	}
	if i >= 6 && i < 9 && j < 3 {
		return 6, 0
	}
	if i >= 6 && i < 9 && j >= 3 && j < 6 {
		return 6, 3
	}
	if i >= 6 && i < 9 && j >= 6 && j < 9 {
		return 6, 6
	}
	return 0, 0
}
*/
