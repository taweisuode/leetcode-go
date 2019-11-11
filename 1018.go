package LeetCode

import (
	"fmt"
	"math"
)

func Code1018() {
	arr := []int{0, 1, 1, 1, 1, 1}
	res := prefixesDivBy5V2(arr)
	fmt.Println(res)
}

/**
给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。

返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。



示例 1：

输入：[0,1,1]
输出：[true,false,false]
解释：
输入数字为 0, 01, 011；也就是十进制中的 0, 1, 3 。只有第一个数可以被 5 整除，因此 answer[0] 为真。
示例 2：

输入：[1,1,1]
输出：[false,false,false]
示例 3：

输入：[0,1,1,1,1,1]
输出：[true,false,false,false,true,false]
示例 4：

输入：[1,1,1,0,1]
输出：[false,false,false,false,false]


提示：

1 <= A.length <= 30000
A[i] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-prefix-divisible-by-5
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, 0)
	if A[0] == 0 {
		res = append(res, true)
	} else {
		res = append(res, false)
	}
	for i := 1; i < len(A); i++ {
		current := getNum(A[:i+1])
		if current%5 == 0 {
			res = append(res, true)
		} else {
			res = append(res, false)
		}

	}
	return res
}

func getNum(a []int) int {
	fmt.Println(a)
	sum := float64(0)
	for i := len(a) - 1; i >= 0; i-- {
		k := math.Pow(2, float64(i))
		sum += k * float64(a[len(a)-i-1])
	}
	fmt.Println(sum)
	return int(sum)
}

func prefixesDivBy5V2(A []int) []bool {
	var temp int
	var res = make([]bool, 0)
	for _, item := range A {
		temp = 2*temp + item
		if temp%5 == 0 {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
		temp %= 5
	}
	return res
}
