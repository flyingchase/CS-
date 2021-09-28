package DataStructure

import (
	"sort"
	"sync"
)

type Set struct {
	sync.RWMutex
	m map[int]bool
}

func (s *Set) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, item := range items {
		s.m[item] = true
	}
}

func New(items ...int) *Set {
	s := &Set{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}

func (s *Set) Duplicate() *Set {
	s.Lock()
	defer s.Unlock()
	r := &Set{
		m: make(map[int]bool, len(s.m)),
	}
	for element := range s.m {
		r.m[element] = true
	}
	return r
}
func (s *Set) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()

	for _, item := range items {
		delete(s.m, item)
	}

}

func (s *Set) Exist(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	//  通过遍历 map 查找是否存在 未找到则 false
	for _, item := range items {
		if _, ok := s.m[item]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Count() int {
	s.Lock()
	defer s.Unlock()
	return len(s.m)
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = make(map[int]bool)
}
func (s *Set) isEmpty() bool {
	s.Lock()
	defer s.Unlock()
	return len(s.m) == 0
}

// 获取 set 内的元素列表 无序
func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 获取 set 内的元素列表有序
func (s *Set) SortedList() []int {
	s.RLock()
	defer s.RUnlock()

	sortedlist := make([]int, 0, len(s.m))
	for item := range s.m {
		sortedlist = append(sortedlist, item)
	}
	sort.Ints(sortedlist)
	return sortedlist
}


// 并集 获取 set 之间的交集存在set内
func (s *Set)Union(sets ...*Set)  {
//	 防止出现死锁 不可锁住多个集合
//	采用副本形式找交集
	r:=s.Duplicate()
	for _,set:=range sets{
		set.Lock()
		for e:=range set.m{
			r.m[e]=true
		}
		set.Unlock()
	}

	s.Lock()
	defer s.Unlock()

	s.m=make(map[int]bool)
	for e:=range r.m{
		s.m[e]=true
	}
}
