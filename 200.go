package LeetCode

import (
	"fmt"
)

func Code200() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

/**
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NodeStr struct {
	Val int
	X   int
	Y   int
}

func numIslands(grid [][]byte) int {
	var dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				queue := &Queue{}
				currentNode := &NodeStr{
					Val: '1',
					X:   i,
					Y:   j,
				}
				inQueue(currentNode, queue)
				for !queueEmpty(queue) {
					popNode := deQueue(queue).(*NodeStr)
					x := popNode.X
					y := popNode.Y
					for k := 0; k < 4; k++ {
						xx := x + dx[k]
						yy := y + dy[k]

						//判断上下左右是否为1
						if xx < 0 || yy < 0 || xx > len(grid)-1 || yy > len(grid[i])-1 {
							continue
						}
						if grid[xx][yy] == '1' {
							node1 := &NodeStr{
								Val: '1',
								X:   xx,
								Y:   yy,
							}
							grid[xx][yy] = '0'
							inQueue(node1, queue)
						}

					}
				}
				count++
			}

		}
	}
	return count
}

func getFirst(grid [][]byte) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return i, j
			}
		}
	}
	return -1, -1
}
