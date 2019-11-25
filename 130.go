package LeetCode

import (
	"fmt"
)

func Code130() {
	board := [][]byte{
		{'X', 'X', 'X', 'O'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}

/**
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。

找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:

X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：

X X X X
X X X X
X X X X
X O X X
解释:

被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surrounded-regions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type BoardNode struct {
	Val    byte
	X      int
	Y      int
	IsUsed int
}

func solve(board [][]byte) {
	boardData := make([][]BoardNode, len(board))
	for i := range board {
		for j := range board[i] {
			node := BoardNode{
				Val:    board[i][j],
				X:      i,
				Y:      j,
				IsUsed: 0,
			}
			boardData[i] = append(boardData[i], node)

		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				bfs1(boardData, i, j)
			}
		}
	}
	fmt.Println(boardData, board)
}

func bfs1(boardData [][]BoardNode, x int, y int) {
	queue := &Queue{}

	node := boardData[x][y]
	inQueue(node, queue)

	if x == 0 || y == 0 || x == len(boardData) || y == len(boardData[0]) {
		return
	}
	for !queueEmpty(queue) {

		size := len(queue.item)

		curInter := deQueue(queue)

		cur := curInter.(BoardNode)
		for i := size; i > 0; i-- {
			//判断四周有没有O
			x := cur.X
			y := cur.Y

			if boardData[x-1][y].Val == 'O' && boardData[x-1][y].IsUsed == 0 {
				boardData[x-1][y].IsUsed = 1
				inQueue(boardData[x-1][y], queue)
			}
			if boardData[x+1][y].Val == 'O' && boardData[x+1][y].IsUsed == 0 {
				boardData[x+1][y].IsUsed = 1
				inQueue(boardData[x+1][y], queue)
			}
			if boardData[x][y-1].Val == 'O' && boardData[x][y-1].IsUsed == 0 {
				boardData[x][y-1].IsUsed = 1
				inQueue(boardData[x][y-1], queue)
			}
			if boardData[x][y+1].Val == 'O' && boardData[x][y+1].IsUsed == 0 {
				boardData[x][y+1].IsUsed = 1
				inQueue(boardData[x][y+1], queue)
			}
		}
	}

	queuePrint(queue)
}
