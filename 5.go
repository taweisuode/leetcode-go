package LeetCode

import (
	"fmt"
)

func Code5() {
	fmt.Println(longestPalindrome("adcdedcaa"))
}

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 *  @desc:  function description
 *  @input: int data
 *  @resp:  interface resp
 *
**/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	//dp := [][]int{}
	dp := make([][]int, len(s))
	start, end := 0, 1
	i, j := 0, 0
	for i = 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i = len(s) - 2; i >= 0; i-- {
		for j = i + 1; j < len(s); j++ {
			if s[i] == s[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
				if dp[i][j] > end {
					start = i
					end = dp[i][j]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return s[start : start+end]
}
