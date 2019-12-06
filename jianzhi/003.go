/**
 * @Time : 2019-11-28 10:32
 * @Author : zhuangjingpeng
 * @File : 003
 * @Desc : file function description
 */
package jianzhi

import "fmt"

var nodeList []int

func Code003() {
	list := InitSingleList([]int{4, 6, 1})
	printListFromTailToHead(list)
	fmt.Println(nodeList)
}

func printListFromTailToHead(head *ListNode) {

	if head != nil {
		printListFromTailToHead(head.Next)
		nodeList = append(nodeList, head.Val)
		fmt.Println(head.Val)
	}
}
