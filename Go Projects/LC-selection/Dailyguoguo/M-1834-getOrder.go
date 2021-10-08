package Dailyguoguo

import (
	"container/heap"
	"sort"
)

type CpuTask struct {
	Index          int
	EnqueueTime    int
	ProcessingTime int
}
type CpuTaskQueue []CpuTask

func (c CpuTaskQueue) Len() int {
	return len(c)
}

func (c CpuTaskQueue) Less(i, j int) bool {
	taskA, taskB := c[i], c[j]
	if taskA.ProcessingTime == taskB.ProcessingTime {
		return taskA.Index < taskB.Index
	}
	return taskA.ProcessingTime < taskB.ProcessingTime
}

func (c CpuTaskQueue) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CpuTaskQueue) Push(x interface{}) {
	*c = append(*c, x.(CpuTask))
}

func (c *CpuTaskQueue) Pop() interface{} {
	old := *c
	x := old[len(old)-1]
	*c = old[:len(old)-1]
	return x
}
func (c CpuTaskQueue) Peek() CpuTask {
	return c[0]
}

func getOrder(tasks [][]int) []int {
	cpuTasks := make([]CpuTask, len(tasks))
	for i, task := range tasks {
		cpuTasks[i] = CpuTask{
			Index:          i,
			EnqueueTime:    task[0],
			ProcessingTime: task[1],
		}
	}
	sort.Slice(cpuTasks, func(i, j int) bool {
		return cpuTasks[i].EnqueueTime < cpuTasks[j].EnqueueTime
	})
	taskQueu := CpuTaskQueue{}

	taskIndex, cpuTime, numProcess := 0, 0, 0
	res := make([]int, len(cpuTasks))
	for taskIndex < len(cpuTasks) || taskQueu.Len() > 0 {
		if taskQueu.Len() == 0 && taskIndex < len(cpuTasks) && cpuTasks[taskIndex].EnqueueTime > cpuTime {
			cpuTime = cpuTasks[taskIndex].EnqueueTime
		}
		for taskIndex < len(cpuTasks) && cpuTasks[taskIndex].EnqueueTime <= cpuTime {
			heap.Push(&taskQueu, cpuTasks[taskIndex])
			taskIndex++
		}
		if taskQueu.Len() > 0 {
			task := heap.Pop(&taskQueu).(CpuTask)
			cpuTime += task.ProcessingTime
			res[numProcess] = task.Index
			numProcess++
		}
	}
	return res
}

func mergeHelper(nums []int, l, mid, r int) {
	if l >= r {
		return
	}
	p1, p2 := l, mid+1
	tmp := make([]int, len(nums))
	i := 0
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			tmp[i] = nums[p1]
			p1++
		} else {
			tmp[i] = nums[p2]
			p2++
		}
		i++
	}
	copy(tmp[i:], nums[p1:mid])
	copy(tmp[i:], nums[p2:r])
	copy(nums, tmp)
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	mergeHelper(nums, l, mid, r)
}
func mergesort(nums []int) {
	if len(nums) == 0 {
		return
	}
	mergeSort(nums, 0, len(nums)-1)
}
