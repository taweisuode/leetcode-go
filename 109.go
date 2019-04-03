package LeetCode

func Code109() {
	head := InitSingleList([]int{-10, -3, 0, 5, 9})

	PrintTree(sortedListToBST(head))
}

/**
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/
func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return createTreeNode2(nums)
	//return createTreeNode(head, nil)
}
func createTreeNode2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = createTreeNode2(nums[0:mid])
	root.Left = createTreeNode2(nums[mid+1:])
	return root
}
func createTreeNode(start *ListNode, end *ListNode) *TreeNode {
	if start == nil || start == end {
		return nil
	}
	fast, slow := start, start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	root := &TreeNode{Val: slow.Val}

	root.Left = createTreeNode(start, slow)
	root.Right = createTreeNode(slow.Next, end)
	return root
}
