package bytedance

import (
	"fmt"
	"sort"
)

func Code1046() {
	intervals := make([]Interval, 0)
	intervals = append(intervals, Interval{
		Start: 1,
		End:   4,
	})
	intervals = append(intervals, Interval{
		Start: 1,
		End:   5,
	})
	/*intervals = append(intervals, Interval{
		Start: 8,
		End:   10,
	})
	intervals = append(intervals, Interval{
		Start: 15,
		End:   18,
	})*/
	fmt.Println(merge(intervals))
}

/**


示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
```
*/

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	start := intervals[0].Start
	end := intervals[0].End

	ans := make([]Interval, 0)

	for _, v := range intervals {
		if v.Start <= end {
			end = Max(end, v.End)
		} else {
			ans = append(ans, Interval{Start: start, End: end})
			start = v.Start
			end = v.End
		}
	}
	ans = append(ans, Interval{Start: start, End: end})
	return ans
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
