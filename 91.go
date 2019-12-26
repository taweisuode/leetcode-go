package LeetCode

import (
	"fmt"
)

func Code91() {
	num := numDecodings("12")
	fmt.Println(num)
}

/**
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numDecodings(s string) int {
	//dp[i] 表示s长度 时候的总数

	lenS := len(s)
	dp := make([]int, lenS+1)
	dp[0], dp[1] = 1, one(s[0:1])
	for i := 2; i <= lenS; i++ {
		w1, w2 := one(s[i-1:i]), two(s[i-2:i])
		dp[i] = dp[i-1]*w1 + dp[i-2]*w2
		if dp[i] == 0 {
			return 0
		}
	}
	fmt.Println(dp)
	return dp[lenS]
}

func one(s string) int {
	if s == "0" {
		return 0
	}
	return 1
}

func two(s string) int {
	if s[0] == '0' {
		return 0
	}
	if s[0] == '1' {
		return 1
	}
	if s[0] == '2' && s[1] <= '6' {
		return 1
	}
	return 0
}
