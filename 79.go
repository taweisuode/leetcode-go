package LeetCode

import (
	"fmt"
)

func Code79() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	boardMap := make([][]int, len(board))
	for i := range board {
		boardMap[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			//判断第一个word 的出现位置
			if board[i][j] == word[0] {
				boardMap[i][j] = 1
				res := dfs_7(board, i, j, word[1:], boardMap)
				if res {
					return true
				} else {
					boardMap[i][j] = 0
				}
			}
		}
	}
	return false
}

func dfs_7(board [][]byte, x int, y int, word string, boardMap [][]int) bool {
	//单词搜索完毕 返回true
	if word == "" {
		return true
	}
	direction := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	for key := range direction {
		nextX := x + direction[key][0]
		nextY := y + direction[key][1]

		if nextX >= 0 && nextX < len(board) && nextY >= 0 && nextY < len(board[0]) {
			if boardMap[nextX][nextY] == 0 && word[0] == board[nextX][nextY] {
				boardMap[nextX][nextY] = 1
				if dfs_7(board, nextX, nextY, word[1:], boardMap) {
					return true
				} else {
					boardMap[nextX][nextY] = 0
					continue
				}
			}
		}
	}
	return false
}
