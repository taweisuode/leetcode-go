package LeetCode

import (
	"encoding/json"
	"fmt"
	"os"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitSingleList(num []int) (listNode *ListNode) {
	headNode := &ListNode{Val: 0, Next: nil}
	listNode = headNode
	for i := 0; i < len(num); i++ {
		newNode := &ListNode{Val: num[i], Next: nil}
		listNode.Next = newNode
		listNode = listNode.Next
	}

	return headNode.Next
}

func PrintSingleList(listNode *ListNode) {
	if listNode == nil {
		fmt.Printf("this link is empty")
	}
	for {
		if listNode != nil {
			fmt.Printf("%d \n", listNode.Val)
			listNode = listNode.Next
		} else {
			goto ENDS
		}
	}
ENDS:
	return
}

type DoubleList struct {
	Val  int
	Next *DoubleList
	Pre  *DoubleList
}

type Queue struct {
	item []interface{}
}

func inQueue(p interface{}, queue *Queue) {
	queue.item = append(queue.item, p)
}

func deQueue(queue *Queue) interface{} {
	length := len(queue.item)
	deleteItem := queue.item[0]
	queue.item = queue.item[1:length]
	return deleteItem
}
func queueEmpty(queue *Queue) bool {
	return len(queue.item) == 0
}
func queuePrint(queue *Queue) {
	if len(queue.item) == 0 {
		fmt.Println("this queue is empty")
	}
	fmt.Println("=====queue start")
	for index, item := range queue.item {
		if index == len(queue.item)-1 {
			fmt.Printf("%+v", item)
		} else {
			fmt.Printf("%+v <-> ", item)
		}
	}
	fmt.Println("\n=====queue end")
}

type Stack struct {
	item []interface{}
}

func StackEmpty(stack *Stack) bool {
	return len(stack.item) == 0
}
func Stackpush(p interface{}, stack *Stack) {
	stack.item = append(stack.item, p)
}
func Stackpop(stack *Stack) interface{} {
	length := len(stack.item)
	p := stack.item[length-1]
	stack.item = stack.item[:length-1]
	return p
}
func stackPrint(stack *Stack) {
	if len(stack.item) == 0 {
		fmt.Println("stack is empty")
	}
	for _, item := range stack.item {
		fmt.Printf("%d\n", item)
	}
}
func StackTop(stack *Stack) interface{} {
	if len(stack.item) == 0 {
		return nil
	}
	p := stack.item[len(stack.item)-1]
	return p
}
func InitTree() (head *TreeNode) {
	head = &TreeNode{}

	return head
}

func PrintTree(head *TreeNode) {
	for _, item := range printTree(head) {
		fmt.Println(item)
	}
}
func printTree(root *TreeNode) [][]int {
	var res [][]int
	queue := &Queue{}
	if root == nil {
		return res
	}
	inQueue(root, queue)
	for !queueEmpty(queue) {
		i := len(queue.item)
		var tempRes []int
		for i > 0 {
			currentNode := deQueue(queue)
			node := currentNode.(*TreeNode)
			tempRes = append(tempRes, node.Val)
			if node.Left != nil {
				inQueue(node.Left, queue)
			}
			if node.Right != nil {
				inQueue(node.Right, queue)

			}
			i--
		}
		res = append(res, tempRes)
	}
	return res
}

func PrintJson(desc string, data interface{}, signal int) {
	res, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(desc, string(res))
	switch signal {
	case 9:
		os.Exit(1)
	default:
		break
	}
}
