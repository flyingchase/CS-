package main 

func sn() {
 
}
func heapsort(nums []int)  {
	if len(nums)==0 {
		return
	}
	size:=len(nums)
	for i,_ := range nums {
		heapinsert(nums,i)
	}
	for size>0 {
		size--
		heapify(nums,0,size)
	}
}

func heapify(nums []int,index int, size int)  {
	left,right:=2*index+1,2*index+2
	var largest int
	for  left<size{
		if right<size&&nums[right]>nums[left] {
			largest=right
		}
		if nums[largest]<nums[[index]] {
		largest=index	
		break
		}
		if largest==index {
			break
		}
		

	}
}
