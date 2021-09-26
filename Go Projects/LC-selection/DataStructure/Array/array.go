package Array

import (
	"errors"
	"fmt"
)

const DefaultCapacity = 16

// capacity最大值
const CapacityThreshold = 1024

type Array struct {
	data     []interface{}
	capacity int
	size     int
}

func New(capacity int) *Array {
	if capacity == 0 {
		return nil

	}
	return &Array{
		capacity: capacity,
		data:     make([]interface{}, capacity, capacity),
		size:     0,
	}
}
func Default() *Array {
	return New(DefaultCapacity)
}

// insert element 注意扩容
func (arr *Array) Add(index int, val interface{}) error {
	err := arr.checkIndex(index)
	if err != nil {
		return err
	}
	//	扩容j
	if arr.size >= arr.capacity {
		if arr.capacity >= CapacityThreshold {
			// 1.25倍扩容 传入的是容量
			arr.resize(int(float64(arr.capacity) * 1.25))
		} else {
			// 2 倍扩容
			arr.resize(arr.capacity << 1)
		}
	}
	// 移出待插入的位置
	copy(arr.data[index+1:], arr.data[index:])
	arr.data[index] = val
	arr.size++
	return nil
}

// 追加 append
func (arr *Array) Append(val interface{}) error {
	// 在数组最后位置追加
	return arr.Add(arr.size, val)
}

// 查找
func (arr *Array) Find(index int) (interface{}, error) {
	err := arr.checkIndex(index)
	if err != nil {
		return nil, err
	}
	return arr.data[index], nil
}

func (arr *Array) checkIndex(index int) error {
	if index < 0 || index > arr.capacity {
		return errors.New("index out of range error")
	}
	return nil
}
func (arr *Array) resize(newCap int) {
	// 创建新的 data 数组再覆盖即可
	data := make([]interface{}, newCap, newCap)
	copy(data, arr.data)
	arr.data = data
	arr.capacity = newCap
}
// 是否包含特定元素 遍历查找
func (arr *Array)Contains(val interface{}) bool  {
	if arr.Empty()  {
		return false
	}
	for _,v:=range arr.data {
		if v==val {
			return true
		}
	}
	return false

}

func (arr *Array) Empty() bool {
	return arr.size==0
}
// 删除 idnex 的值
func (arr *Array) Delete (index int) error {
	err:=arr.checkIndex(index)
	if err != nil {
		return err
	}
	if arr.data[index]==nil {
		return nil
	}
	// index 下表位置置空并前移数组
	arr.data[index]=nil
	copy(arr.data[index:],arr.data[index+1:])
	// 处理 size 和数组的最后一位
	arr.data[arr.size-1]=nil
	arr.size--
	return nil
}
func (arr *Array)Size() int  {
	return arr.size
}
func (arr *Array)PrintData()  {
	for i:=0;i<arr.capacity;i++ {
		if arr.data[i]!=nil {
			fmt.Print(arr.data[i]," ")
		}
	}
	fmt.Println()
}