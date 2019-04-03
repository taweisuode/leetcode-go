package bytedance

import (
	"math/rand"
)

func MakeRand() chan int {
	res := make(chan int, 0)
	go func() {
		for {
			res <- rand.Int()
		}

	}()

	go func() {
		for {
			res <- rand.Int()
		}
	}()
	return res
}

/**

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
```
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
