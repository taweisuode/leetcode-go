package bytedance

import (
	"fmt"
	"strconv"
)

func Code1015() {
	s1 := "1231312312313214314242424324234131231312312312312312312312312414324234234234234234324"
	s2 := "456"
	fmt.Println(multiply(s1, s2))
}

/**

给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func multiply(num1 string, num2 string) string {
	int1, _ := strconv.ParseUint(num1, 10, 64)
	int2, _ := strconv.ParseUint(num2, 10, 64)
	fmt.Println(int1)
	fmt.Println(int2)
	return ""
}
