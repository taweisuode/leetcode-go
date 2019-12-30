/**
 * @Time : 2019-12-30 17:34
 * @Author : zhuangjingpeng
 * @File : zhishu
 * @Desc : file function description
 */
package bytedance

import (
	"fmt"
)

func CodeZhishu() {
	fmt.Println(getPrimeList(20))
}

/**

求 1-n  以内的  所有素数
```
*/

//挨式筛法
func getPrimeList(num int) []int {
	var res []int
	resMap := make(map[int]bool, 0)
	for i := 2; i < num; i++ {
		if !resMap[i] {
			res = append(res, i)
			for j := 2 * i; j < num; j = j + i {
				resMap[j] = true
			}
		}
	}
	return res
}
