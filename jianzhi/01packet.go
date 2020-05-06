/**
 * @Time : 2020-01-17 11:27
 * @Author : zhuangjingpeng
 * @File : 01packet
 * @Desc : file function description
 */
package jianzhi

import "fmt"

func Code01Packet() {

	productMap := map[int]int{
		5: 12,
		4: 3,
		7: 10,
		2: 3,
		6: 6,
	}

	fmt.Println(maxProfit(productMap, 10))
}

func maxProfit(productMap map[int]int, load int) int {

}
