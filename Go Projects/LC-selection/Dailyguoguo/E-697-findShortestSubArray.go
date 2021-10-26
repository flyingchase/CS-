package Dailyguoguo

// double pointers extend and shrink
func findShortestSubArray(nums []int) int {
	size, left, right, degree := len(nums), 0, 0, 0

	res := size
	mapR := make(map[int]int, 0)
	for _, num := range nums {
		mapR[num]++
		degree = max697(mapR[num], degree)
	}

	countMap := make(map[int]int, 0)
	for right < size {
		for right < size {
			countMap[nums[right]]++
			right++
			if countMap[nums[right-1]] == degree {
				break
			}
		}
		for left < right {
			countMap[nums[left]]--
			left++
			if countMap[nums[left-1]] == degree-1 {
				res = min697(res, right-left+1)
				break
			}
		}
	}
	return res
}

func min697(j int, i int) int {
	if i < j {
		return i
	}
	return j
}

func max697(i int, degree int) int {
	if i < degree {
		return degree
	}
	return i
}
