package LeetCode

import (
	"fmt"
	"strconv"
)

func Code62() {
	fmt.Println(uniquePaths(7, 3))
}

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，上图是一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func uniquePaths(m int, n int) int {

	return bfs2(m, n)
}
func uniquePaths2(m int, n int) int {

	dfsMap := make(map[string]int, 0)
	return dfs2(m, n, 1, 1, dfsMap)

}

type nodeStr struct {
	X      int
	Y      int
	IsUsed int
}

//curM 为当前所在的m  curN 为当前所在的n
func dfs2(m int, n int, curM int, curN int, dfsMap map[string]int) int {
	if dfsMap[strconv.Itoa(curM)+"_"+strconv.Itoa(curN)] > 0 {
		return dfsMap[strconv.Itoa(curM)+"_"+strconv.Itoa(curN)]
	}
	if curM == m && curN == n {
		return 1
	}

	sum := 0

	if curM < m {
		sum += dfs2(m, n, curM+1, curN, dfsMap)
	}

	if curN < n {
		sum += dfs2(m, n, curM, curN+1, dfsMap)
	}

	dfsMap[strconv.Itoa(curM)+"_"+strconv.Itoa(curN)] = sum
	return sum
}

func bfs2(m int, n int) int {
	firstNode := nodeStr{
		X: 0,
		Y: 0,
	}
	nodeMap := make(map[nodeStr]int)
	queue := &Queue{}
	inQueue(firstNode, queue)

	count := -1
	for !queueEmpty(queue) {
		size := len(queue.item)
		count++
		cur := deQueue(queue)
		curNode := cur.(nodeStr)

		for i := size; i > 0; i-- {
			list := queue.getQueueList(curNode, m, n, nodeMap)

			fmt.Println(list)
			for index := range list {
				inQueue(list[index], queue)
			}
		}
	}
	return count
}

func (queue *Queue) getQueueList(node nodeStr, m int, n int, nodeMap map[nodeStr]int) []nodeStr {
	var nodeList []nodeStr
	if node.Y < n-1 {
		rightNode := nodeStr{
			X:      node.X,
			Y:      node.Y + 1,
			IsUsed: 1,
		}
		if _, ok := nodeMap[rightNode]; !ok {
			nodeList = append(nodeList, rightNode)
			nodeMap[rightNode] = 1
		}

	}
	if node.X < m-1 {
		downNode := nodeStr{
			X:      node.X + 1,
			Y:      node.Y,
			IsUsed: 1,
		}
		if _, ok := nodeMap[downNode]; !ok {
			nodeList = append(nodeList, downNode)
			nodeMap[downNode] = 1
		}
	}
	return nodeList
}
