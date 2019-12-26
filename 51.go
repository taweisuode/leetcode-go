package LeetCode

import (
	"fmt"
	"math"
)

func Code51() {
	fmt.Println(solveNQueens(4))
}

/**
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例:

输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func solveNQueens1(n int) [][]string {
	//声明返回结果
	var result [][]string

	//判断长度
	if n <= 0 {
		return result
	}
	if n == 1 {
		result = append(result, []string{"Q"})
		return result
	}

	//---根据皇后的攻击特性可以确定N个皇后，一行只能存在1个皇后
	var posResult [][]int
	var queens []int

	//回溯
	f6(queens, 0, n, &posResult)

	fmt.Println(posResult)

	//处理结果集
	for _, queen := range posResult {
		//queen为每一行皇后排放的位置
		var tmpStrArr []string
		for i := 0; i < n; i++ { //行数
			//每一行的列
			tmpStr := ""
			for j := 0; j < n; j++ {
				//判断是否为皇后
				if j == queen[i] {
					tmpStr += "Q"
				} else {
					tmpStr += "."
				}
			}

			tmpStrArr = append(tmpStrArr, tmpStr)

		}

		//存放当前结果集
		result = append(result, tmpStrArr)

	}
	//返回结果集
	return result
}

/**
@param queens[]int 已经放置的皇后
@param current int 当前的深度
@param height int 最大深度（皇后的个数）
@param result *[][]int 存放所有的结果
*/
func f6(queens []int, current, height int, result *[][]int) {
	//判断当前深度
	if current == height {
		//存放结果
		*result = append(*result, queens)
		return
	}

	//抉择每一行
	for i := 0; i < height; i++ {
		//判断当前是否可以放置皇后
		if f6isOk(queens, i) {
			//深拷贝之前放置的皇后
			queens1 := make([]int, len(queens))
			copy(queens1, queens)

			//存放当前的皇后
			queens1 = append(queens1, i)

			//放置下一个
			f6(queens1, current+1, height, result)

		}

	}
}

/**
判断是否可以放置皇后
@param queens []int  之前放置的
@param current 当前要放置的位置
*/
func f6isOk(queens []int, current int) bool {
	//声明返回
	var isOk = true

	//判断是否为空
	if len(queens) == 0 {
		return true
	}

	//判断是否在同一列--queens 有相等的值
	for _, v := range queens {
		//判断同一列
		if v == current {
			isOk = false
			break
		}
	}

	//判断
	if !isOk {
		return isOk
	}

	//当前的位置
	pos := len(queens)

	//判断对角线
	for k, v := range queens {
		//第k位置
		posL, posR := v, v //对角线左，对角线右
		for j := k + 1; j <= pos; j++ {
			//左对角线
			posL--

			//右对角线
			posR++
		}

		//判断是否攻击
		if posL == current || posR == current {
			isOk = false
		}

	}
	//返回
	return isOk
}

var n int

// 用于统计第row行的queen在第几列
var nums []int

func solveNQueens(n1 int) [][]string {
	n = n1
	nums = make([]int, n)
	var result [][]string
	backtrack(0, &result)

	return result
}

// 判断放置的queen在当前位置是否有效
func isValid1(nums []int, row int) bool {
	for i := 0; i < row; i++ {
		if math.Abs(float64(nums[row]-nums[i])) == math.Abs(float64(row-i)) || nums[i] == nums[row] {
			return false
		}
	}
	return true
}

func addSolution(nums []int, result *[][]string) {
	// fmt.Println(nums)
	var x []string
	for i := 0; i < len(nums); i++ {
		var s string
		for j := 0; j < n; j++ {
			if j == nums[i] {
				s += "Q"
			} else {
				s += "."
			}
		}

		x = append(x, s)
	}
	*result = append(*result, x)
}

// 开始进行回溯
func backtrack(row int, result *[][]string) {
	if row == n {
		addSolution(nums, result)
		return
	} else {
		for col := 0; col < n; col++ {
			nums[row] = col
			if isValid1(nums, row) {
				backtrack(row+1, result)
			}
		}
	}
}
