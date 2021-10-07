package WarmUp

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	lists := make([]int, 0)
	dfsSubset2(&res, lists, 0, nums)
	return res
}

func dfsSubset2(res *[][]int, lists []int, index int, nums []int) {
	// tmp:=make([]int,len(lists))
	// copy(tmp,lists)
	// *res=append(*res,tmp)
	*res = append(*res, append([]int{}, lists...))
	// 避免三句代码 使用 temp
	for i := index; i < len(nums); i++ {
		lists = append(lists, nums[i])
	}
	if i > index && nums[i] == nums[i-1] {
		continue
	}
	dfsSubset2(res, lists, i+1, nums)
	lists = lists[:len(lists)-1]
}
