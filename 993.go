package LeetCode

import (
	"fmt"
)

func Code993() {
	head := InitTree()
	head = &TreeNode{1, nil, nil}
	head.Left = &TreeNode{2, nil, nil}
	head.Right = &TreeNode{3, nil, nil}
	head.Left.Right = &TreeNode{4, nil, nil}
	head.Right.Right = &TreeNode{5, nil, nil}

	fmt.Println(isCousins(head, 5, 4))
}

/**
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。



示例 1：


输入：root = [1,2,3,4], x = 4, y = 3
输出：false
示例 2：


输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
输出：true
示例 3：



输入：root = [1,2,3,null,4], x = 2, y = 3
输出：false


提示：

二叉树的节点数介于 2 到 100 之间。
每个节点的值都是唯一的、范围为 1 到 100 的整数。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cousins-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	return false
}
