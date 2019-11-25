package LeetCode

import (
	"fmt"
)

func Code389() {
	s := "abcd"
	t := "abcde"
	num := findTheDifference(s, t)
	fmt.Println(num)
}

/**
给定两个字符串 s 和 t，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。



示例:

输入：
s = "abcd"
t = "abcde"

输出：
e

解释：
'e' 是那个被添加的字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-difference
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findTheDifference(s string, t string) byte {
	var res byte
	for _, item := range []byte(s) {
		res ^= item
	}
	for _, item2 := range []byte(t) {
		res ^= item2
	}
	return res
}
