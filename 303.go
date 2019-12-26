package LeetCode

import (
	"fmt"
)

func Code303() {
	NumArray := Constructor1([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(NumArray.SumRange(0, 2))
}

/**
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：

给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type NumArray struct {
	arr []int
}

func Constructor1(nums []int) NumArray {
	arr := make([]int, len(nums))
	numArr := NumArray{
		arr: arr,
	}
	for i := range nums {
		numArr.arr[i] = nums[i]
	}
	return numArr
}

func (this *NumArray) SumRange(i int, j int) int {
	sum := 0
	for x := i; x <= j; x++ {
		sum += this.arr[x]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
