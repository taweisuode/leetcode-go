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
	boardMapI := make(map[int]int, len(board))
	boardMapJ := make(map[int]int, len(board[0]))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			boardMapI[i] = 0
			boardMapJ[j] = 0
		}
	}

	fmt.Println(boardMapI, boardMapJ)
	for i := 0; i < len(word); i++ {
		getNearNode(board, 'A', boardMapI, boardMapJ)
	}
	return false
}

func getNearNode(board [][]byte, node byte, boardMapI map[int]int, boardMapJ map[int]int) [][]byte {
	res := [][]byte{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == node {

				//上边
				if i > 0 && boardMapI[i] != 1 {
					res = append(res, []byte{board[i-1][j]})
				}

				//左边
				if j > 0 && boardMapJ[j] != 1 {
					res = append(res, []byte{board[i][j-1]})
				}

				//下边
				if i < len(board)-1 && boardMapI[i] != 1 {
					fmt.Println(i)
					res = append(res, []byte{board[i+1][j]})
				}

				//右边
				if j < len(board[0])-1 && boardMapJ[j] != 1 {
					res = append(res, []byte{board[i][j+1]})
				}
			}
		}
	}
	fmt.Println(res)
	return nil
}
