package LeetCode

import (
	"fmt"
)

func Code139() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak3(s, wordDict))
}

/**
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {

		for j := i + 1; j < len(s)+1; j++ {
			if dp[i] && checkInArr(s[i:j], wordDict) {
				dp[j] = true
			}
		}
	}

	fmt.Println(dp)
	return dp[len(dp)-1]
}

func checkInArr(s string, arr []string) bool {
	for i := range arr {
		if arr[i] == s {
			return true
		}
	}

	return false
}

func wordBreak2(s string, wordDict []string) bool {
	return dfs_139(s, 0, wordDict)
}

func dfs_139(s string, start int, wordDict []string) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		if checkInArr(s[start:i], wordDict) && dfs_139(s, i, wordDict) {
			return true
		}
	}

	return false
}

func wordBreak3(s string, wordDict []string) bool {

	queue := &Queue{}
	inQueue(0, queue)

	visit := make([]int, len(s))
	for !queueEmpty(queue) {
		cur := deQueue(queue).(int)

		if visit[cur] == 0 {
			for i := cur + 1; i <= len(s); i++ {
				if checkInArr(s[cur:i], wordDict) {
					inQueue(i, queue)
					if i == len(s) {
						return true
					}
				}
			}
			visit[cur] = 1
		}
	}
	return false
}
