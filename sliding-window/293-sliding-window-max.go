package main

import (
	"container/heap"
	"fmt"
)

// 293 max sliding window: using hash heap, refer utils/treemap.go
type SWHeap struct {
	ele 	 []*Ele
	hash     map[int]int
}

type Ele struct {
	value int
	index int
}

func (s *SWHeap) Swap(i, j int) {
	s.hash[s.ele[i].index], s.hash[s.ele[j].index] = s.hash[s.ele[j].index], s.hash[s.ele[i].index]
	s.ele[i], s.ele[j] = s.ele[j], s.ele[i]
}

func (s *SWHeap) Len() int {
	return len(s.ele)
}

func (s *SWHeap) Less(i, j int) bool {
	if s.ele[i].value > s.ele[j].value {
		return true
	}
	return false
}

func (s *SWHeap) Push(ele interface{}) {
	(*s).ele = append((*s).ele, ele.(*Ele))
	(*s).hash[ele.(*Ele).index] = (*s).Len() - 1
}

func (s *SWHeap) Pop() interface {} {
	ele := (*s).ele[len((*s).ele) - 1]
	delete((*s).hash, s.ele[len((*s).ele) - 1].index)
	(*s).ele = (*s).ele[: len((*s).ele) - 1]
	return ele
}

func (s *SWHeap) Front() interface{} {
	ele := heap.Pop(s)
	heap.Push(s, ele)
	return ele
}

func (s *SWHeap) Remove(index int) {
	if val, exist := s.hash[index]; exist {
		heap.Remove(s, val)
	} else {
		fmt.Println("miss")
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	results := make([]int, 0)
	swheap := &SWHeap{
		ele: 	  make([]*Ele, 0),
		hash:     map[int]int{},
	}
	heap.Init(swheap)
	for i := 0; i < k; i ++ {
		heap.Push(swheap, &Ele{nums[i], i})
	}
	results = append(results, swheap.Front().(*Ele).value)
	for i := k; i < len(nums); i ++ {
		swheap.Remove(i-k)
		heap.Push(swheap, &Ele{nums[i], i})
		if swheap.Len() != k {
			panic("heap length not match... ")
		}
		results = append(results, swheap.Front().(*Ele).value)
	}
	return results
}
