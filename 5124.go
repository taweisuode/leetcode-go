package LeetCode

import (
	"fmt"
	"strconv"
)

func Code5124() {
	fmt.Println(sequentialDigits(1000, 13000))
}

/**
我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。

请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。



示例 1：

输出：low = 100, high = 300
输出：[123,234]
示例 2：

输出：low = 1000, high = 13000
输出：[1234,2345,3456,4567,5678,6789,12345]
*/
func sequentialDigits(low int, high int) []int {
	var res []int
	for i := low; i <= high; i++ {
		if judge(i) {
			res = append(res, i)
		}
	}
	return res
}
func judge(num int) bool {
	each := strconv.Itoa(num)
	count := 0
	for x := 1; x < len(each); x++ {
		if each[x] == each[x-1]+1 {
			count++
		} else {
			return false
		}
	}
	if count == len(each)-1 {
		return true
	}
	return false
}
