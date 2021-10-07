package Dailyguoguo

import "sort"

// 所有非重复数字中第三大的数字

func thirdMax(nums []int) int {
	set := make(map[int]bool)
	windows := make([]int, 0, 4)
	for _, num := range nums {
		if _, ok := set[num]; ok {
			continue
		}
		set[num] = true
		windows = append(windows, num)
		sort.Ints(windows)
		if len(indows) > 3 {
			windows = windows[1:]
		}
	}
	if len(windows) < 3 {
		return windows[len(windows)-1]
	}
	return windows[0]
}
