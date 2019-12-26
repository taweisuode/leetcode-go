package LeetCode

import "fmt"

func Code54() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	fmt.Println(spiralOrder(matrix))
}

/**
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type matrixNode struct {
	Val int
	X   int
	Y   int
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	matrixMap := make(map[matrixNode]int)
	for x := range matrix {
		for y := range matrix[x] {
			node := matrixNode{
				Val: matrix[x][y],
				X:   x,
				Y:   y,
			}
			matrixMap[node] = 0
		}
	}
	L := len(matrix)
	//H := len(matrix[0])

	for x := range matrix {
		for y := range matrix[x] {
			//节点先判断右边有没有未使用节点
			if y+1 < L {
				node := matrixNode{
					Val: matrix[x][y+1],
					X:   x,
					Y:   y + 1,
				}
				fmt.Println(x, y, matrix[x][y], node, matrixMap[node])
				if matrixMap[node] == 0 {
					res = append(res, matrix[x][y])
					matrixMap[node] = 1
				}
			}
		}
	}

	fmt.Println(res)
	return nil
}
