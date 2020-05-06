package LeetCode

import (
	"fmt"
	"os"
	"strings"
)

func Code20() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isValid(s string) bool {
	s1 := "hello zhuang jing peng"

	arr := strings.Split(s1, " ")
	res := ""
	for _, item := range arr {
		res += RR(item) + " "
		//res = RR(item) + " "
	}
	res = RR(res)
	fmt.Println(res)
	os.Exit(1)
	stackMap := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := &Stack{}
	if len(s) <= 0 {
		return true
	}
	Stackpush(s[0], stack)
	for i := 1; i < len(s); i++ {
		var stackTop byte
		if !StackEmpty(stack) {
			stackTop = StackTop(stack).(byte)
		}
		if item, ok := stackMap[s[i]]; ok {
			if item == stackTop {
				Stackpop(stack)
			} else {
				Stackpush(s[i], stack)
			}
		} else {
			Stackpush(s[i], stack)
		}
	}

	fmt.Println("after")
	stackPrint(stack)
	return StackEmpty(stack)
}

func RR(s string) string {
	runes := []rune(s)

	for start, end := 0, len(runes)-1; start < end; {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}
