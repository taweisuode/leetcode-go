package LeetCode

import (
	"fmt"
	"strconv"
)

func Code150() {
	token := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(token))
}

/**
根据逆波兰表示法，求表达式的值。

有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：

输入: ["2", "1", "+", "3", "*"]
输出: 9
解释: ((2 + 1) * 3) = 9
示例 2：

输入: ["4", "13", "5", "/", "+"]
输出: 6
解释: (4 + (13 / 5)) = 6
示例 3：

输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
输出: 22
解释:
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func evalRPN(tokens []string) int {
	stack := &Stack1{}
	operationMap := map[string]int{
		"+": 1,
		"-": 1,
		"*": 1,
		"/": 1,
	}
	for i := 0; i < len(tokens); i++ {
		if _, ok := operationMap[tokens[i]]; ok {

			secondNode := Stackpop1(stack)
			firstNode := Stackpop1(stack)
			res := cal(firstNode, secondNode, tokens[i])
			Stackpush1(res, stack)
		} else {
			each, _ := strconv.Atoi(tokens[i])
			Stackpush1(each, stack)
		}
	}
	newRes := Stackpop1(stack)
	return newRes

}
func cal(one int, two int, operationMap string) int {
	switch operationMap {
	case "+":
		return one + two
	case "-":
		return one - two
	case "*":
		return one * two
	case "/":
		return one / two
	}
	return 0
}

type Stack1 struct {
	item []int
}

func StackEmpty1(stack *Stack1) bool {
	return len(stack.item) == 0
}
func Stackpush1(p int, stack *Stack1) {
	stack.item = append(stack.item, p)
}
func Stackpop1(stack *Stack1) int {
	length := len(stack.item)
	p := stack.item[length-1]
	stack.item = stack.item[:length-1]
	return p
}
