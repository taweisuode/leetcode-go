package LeetCode

import (
	"fmt"
	"math"
)

//love 函数  x^2 + (y-(x^2)^1/3)^2 = 1
func LoveStar() {
	word := "love"

	allStr := make([]string, 0)

	for y := -13; y < 13; y++ {
		line := make([]string, 0)
		lineCon := ""
		for x := -30; x < 30; x++ {
			xFloat := float64(x)
			yFloat := float64(y)
			loc := math.Pow(math.Pow(xFloat*0.054, 2)+math.Pow(yFloat*0.15, 2)-1, 3) + math.Pow(xFloat*0.2, 2)*math.Pow(yFloat*0.1, 3)
			//fmt.Println(xFloat, yFloat, loc)
			if loc <= 0 {
				index := int(x) % len(word)
				if index >= 0 {
					lineCon += string(word[index])
				} else {
					lineCon += string(word[int(float64(len(word))-math.Abs(float64(index)))])
				}
			} else {
				lineCon += " "
			}

		}
		line = append(line, lineCon)
		allStr = append(allStr, line...)
	}

	for _, text := range allStr {
		fmt.Printf("%s\n", text)
	}
}
