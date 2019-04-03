package LeetCode

import (
	"fmt"
	"math"
)

func Code6() {
	fmt.Println(convert("LEETCODEISHIRING", 3))
}

/**
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	//只有当我们向上移动到最上面或向下移动到最下面的行时，当前的方向才会发生改变
	//计算出需要的行数
	bLen := int(math.Min(float64(numRows), float64(len(s))))
	bytes := make([][]byte, bLen)
	//设置控制移动的标志变量
	goingDownFlag := false
	curRow := 0
	//遍历字符串进行每个行的值的添加
	for _, v := range s {
		bytes[curRow] = append(bytes[curRow], byte(v))
		if curRow == 0 || curRow == numRows-1 {
			goingDownFlag = !goingDownFlag
		}

		if goingDownFlag {
			curRow++
		} else {
			curRow--
		}
	}

	//遍历并输出最后的结果
	result := ""
	for _, v := range bytes {
		result += string(v)
	}
	return result
}
