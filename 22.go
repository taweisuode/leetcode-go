package LeetCode

import "fmt"

func Code22() {
	fmt.Println(generateParenthesis(3))
}

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func generateParenthesis(n int) []string {
	var res []string
	left, right := n, n
	dfs_22(n, left, right, &res, "")
	return res
}

func dfs_22(n int, left int, right int, res *[]string, str string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}

	//左边还有剩余时
	if left > 0 {
		dfs_22(n, left-1, right, res, str+"(")
	}

	//右边有剩余并且 左边剩余比右边少
	if right > 0 && left < right {
		dfs_22(n, left, right-1, res, str+")")
	}
}
