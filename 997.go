package LeetCode

import (
	"fmt"
)

func Code997() {
	trust := [][]int{
		{1, 2}, {2, 3},
	}
	fmt.Println(findJudge(3, trust))
}

/**
在一个小镇里，按从 1 到 N 标记了 N 个人。传言称，这些人中有一个是小镇上的秘密法官。

如果小镇的法官真的存在，那么：

小镇的法官不相信任何人。
每个人（除了小镇法官外）都信任小镇的法官。
只有一个人同时满足属性 1 和属性 2 。
给定数组 trust，该数组由信任对 trust[i] = [a, b] 组成，表示标记为 a 的人信任标记为 b 的人。

如果小镇存在秘密法官并且可以确定他的身份，请返回该法官的标记。否则，返回 -1。



示例 1：

输入：N = 2, trust = [[1,2]]
输出：2
示例 2：

输入：N = 3, trust = [[1,3],[2,3]]
输出：3
示例 3：

输入：N = 3, trust = [[1,3],[2,3],[3,1]]
输出：-1
示例 4：

输入：N = 3, trust = [[1,2],[2,3]]
输出：-1
示例 5：

输入：N = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
输出：3


提示：

1 <= N <= 1000
trust.length <= 10000
trust[i] 是完全不同的
trust[i][0] != trust[i][1]
1 <= trust[i][0], trust[i][1] <= N

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-town-judge
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findJudge(N int, trust [][]int) int {
	inMap := make(map[int]int, 0)
	outMap := make(map[int]int, 0)

	for i := 0; i < len(trust); i++ {
		if _, ok := inMap[trust[i][0]]; !ok {
			inMap[trust[i][0]] = 1
		} else {
			inMap[trust[i][0]]++
		}

		if _, ok := outMap[trust[i][1]]; !ok {
			outMap[trust[i][1]] = 1
		} else {
			outMap[trust[i][1]]++
		}
	}

	for i := range outMap {
		if _, ok := inMap[i]; !ok && outMap[i] == N-1 {
			return i
		}
	}
	return -1
}
