package LeetCode

import (
	"fmt"
)

func Code230() {
	headTree := InitTree()
	headTree = &TreeNode{5, nil, nil}
	headTree.Left = &TreeNode{3, nil, nil}
	headTree.Right = &TreeNode{6, nil, nil}
	headTree.Left.Left = &TreeNode{2, nil, nil}
	headTree.Left.Right = &TreeNode{4, nil, nil}
	headTree.Left.Left.Left = &TreeNode{1, nil, nil}
	fmt.Println(kthSmallest(headTree, 3))
}

/**
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {

	treeList := make([]int, 0)
	list := preReverse(root, treeList)
	fmt.Println(list)
	//sort.Ints(list)
	return list[k-1]
}

func preReverse(root *TreeNode, list []int) []int {
	if root != nil {
		list = preReverse(root.Left, list)
		fmt.Println(list)
		list = append(list, root.Val)
		list = preReverse(root.Right, list)
	}
	return list
}
