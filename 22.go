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
	//获取所有的组合
	getFullData(n)
	return []string{}
}

func getFullData(n int) {

}
