package LeetCode

import (
	"fmt"
)

func Code863() {
	head := InitTree()
	/*head = &TreeNode{3, nil, nil}
	head.Left = &TreeNode{5, nil, nil}
	head.Right = &TreeNode{1, nil, nil}
	head.Left.Left = &TreeNode{6, nil, nil}
	head.Left.Right = &TreeNode{2, nil, nil}
	head.Right.Left = &TreeNode{0, nil, nil}
	head.Right.Right = &TreeNode{8, nil, nil}
	head.Left.Right.Left = &TreeNode{7, nil, nil}
	head.Left.Right.Left.Left = &TreeNode{11, nil, nil}
	head.Left.Right.Right = &TreeNode{4, nil, nil}
	head.Left.Right.Right.Left = &TreeNode{13, nil, nil}
	target := head.Left*/
	head = &TreeNode{0, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{1, nil, nil}
	head.Left.Left = &TreeNode{2, nil, nil}
	target := head.Left.Left
	fmt.Println(distanceK(head, target, 3))
}

/**
输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2

输出：[7,4,1]

解释：
所求结点为与目标结点（值为 5）距离为 2 的结点，
值分别为 7，4，以及 1



注意，输入的 "root" 和 "target" 实际上是树上的结点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	queue := &Queue{}
	nodeMap := make(map[*TreeNode]int)
	graph := make(map[*TreeNode][]*TreeNode)
	graphMap := buildGraph(root, graph)
	//printMap(graphMap)
	inQueue(target, queue)
	var count int
	for !queueEmpty(queue) {
		var level []int
		size := len(queue.item)
		for i := size; i > 0; i-- {
			popNode := deQueue(queue).(*TreeNode)
			if _, ok := nodeMap[popNode]; ok {
				break
			}
			for _, item := range graphMap[popNode] {
				if _, ok := nodeMap[item]; !ok && item != target {
					inQueue(item, queue)
					nodeMap[popNode] = 1
					level = append(level, item.Val)
				}
			}
		}
		count++
		if count == K {
			return level
		}
	}
	return nil
}

func buildGraph(root *TreeNode, graph map[*TreeNode][]*TreeNode) map[*TreeNode][]*TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		graph[root] = append(graph[root], root.Left)
		graph[root.Left] = append(graph[root.Left], root)
		buildGraph(root.Left, graph)
	}
	if root.Right != nil {
		graph[root] = append(graph[root], root.Right)
		graph[root.Right] = append(graph[root.Right], root)
		buildGraph(root.Right, graph)
	}
	return graph
}

func printMap(graph map[*TreeNode][]*TreeNode) {

	for index, item := range graph {
		fmt.Println(index.Val)
		var level []int
		for i := range item {
			level = append(level, item[i].Val)
		}
		fmt.Println(level)
	}
}
