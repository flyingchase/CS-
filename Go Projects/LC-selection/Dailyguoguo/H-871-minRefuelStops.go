package Dailyguoguo

import "container/heap"

type maxHeap871 []int

func (h maxHeap871) Len() int {
	return len(h)
}
func (h maxHeap871) Swap(i, j int) {

	h[i], h[j] = h[j], h[i]
}
func (h maxHeap871) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *maxHeap871) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap871) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {

	h, pos, count := new(maxHeap871), 0, 0
	tryFulUplRetrospectively := func(distance int) bool {
		fuel := distance
		for fuel < 0 && h.Len() > 0 {
			fuel += heap.Pop(h).(int)
			count++
		}
		return fuel >= 0
	}
	for _, station := range stations {
		if !tryFulUplRetrospectively(station[0] - pos) {
			return -1
		}
		heap.Push(h, station[1])
		pos = station[0]
	}
	if !tryFulUplRetrospectively(target - pos) {
		return -1
	}

	return count

}
