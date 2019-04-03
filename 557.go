package LeetCode

import (
	"fmt"
	"strings"
)

func Code557() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseWords(s string) string {
	res := ""
	arr := strings.Split(s, " ")
	for _, item := range arr {
		data := reverse(item)
		res += data + " "
	}
	return res[0 : len(res)-1]
}
func reverse(s string) string {
	str := []byte(s)
	for i, j := 0, len(str)-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}
