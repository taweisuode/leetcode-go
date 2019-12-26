package LeetCode

import (
	"fmt"
)

func Code516() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

/**
给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。

示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func longestPalindromeSubseq(s string) int {
	if s == "" {
		return 0
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
	return len(s[start : start+end])
}
