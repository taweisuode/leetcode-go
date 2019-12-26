package LeetCode

import (
	"fmt"
)

func Code264() {
	num := nthUglyNumber(100)
	fmt.Println(num)
}

/**
编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ugly-number-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * 超时解法1 会超时..
 *
**/
/*func nthUglyNumber(n int) int {
	pool := make([]int, 0)
	if n == 1 {
		return 1
	}
	for i := 1; i < math.MaxInt64; i++ {
		temp := i
		if len(pool) == n {
			break
		}
		for temp%2 == 0 {
			temp = temp / 2
		}
		for temp%3 == 0 {
			temp = temp / 3
		}
		for temp%5 == 0 {
			temp = temp / 5
		}
		if temp == 1 {
			pool = append(pool, i)
		}
	}
	return pool[len(pool)-1]
}*/

func nthUglyNumber(n int) int {
	if n == 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	i, int2, int3, int5 := 1, 0, 0, 0
	stack := &Stack{}
	for i < n {
		a := 2 * dp[int2]
		b := 3 * dp[int3]
		c := 5 * dp[int5]
		dp[i] = minThree(2*dp[int2], 3*dp[int3], 5*dp[int5])
		Stackpush(dp[i], stack)
		if dp[i] == a {
			int2++
		}
		if dp[i] == b {
			int3++
		}
		if dp[i] == c {
			int5++
		}
		i++
	}
	return Stackpop(stack).(int)
}

func minThree(a, b, c int) int {
	temp := a
	if a > b {
		temp = b
		if temp > c {
			return c
		} else {
			return temp
		}
	} else {
		if temp > c {
			return c
		} else {
			return temp
		}
	}

}
