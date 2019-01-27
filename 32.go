package LeetCode

import (
	"fmt"
)

func Code32() {
	fmt.Println(longestValidParentheses("()(()"))
}

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func longestValidParentheses(s string) int {
	stack := &Stack{nil}
	nodeMap := map[string]string{}
	nodeMap["("] = ")"
	if len(s) == 0 {
		return 0
	}
	num := 0
	for i := 0; i < len(s); i++ {
		if len(stack.item) == 0 {
			Stackpush(string(s[i]), stack)
		} else {
			if char, ok := StackTop(stack).(string); ok {
				if nodeMap[char] != string(s[i]) {
					Stackpush(string(s[i]), stack)
				} else {
					Stackpop(stack)
					num = num + 2
				}
			}
		}
		fmt.Println(stack.item)
	}
	return num
}
