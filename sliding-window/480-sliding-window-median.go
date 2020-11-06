package main

import (
	"container/heap"
)

// 480 sliding window median using hash heap
type intHeap struct {
	arr     []*Element
	hash	map[int]int
	isSmall bool
}

type Element struct {
	value int
	index int
}

func (ih *intHeap) Len() int {
	return len((*ih).arr)
}

func (ih *intHeap) Less(i, j int) bool {
	if ih.isSmall {
		return (*ih).arr[i].value < (*ih).arr[j].value
	} else {
		return (*ih).arr[i].value > (*ih).arr[j].value
	}
}

func (ih *intHeap) Push(ele interface{}) {
	(*ih).arr = append((*ih).arr, ele.(*Element))
	(*ih).hash[ele.(*Element).index] = (*ih).Len() - 1
}

func (ih *intHeap) Swap(i, j int) {
	(*ih).hash[(*ih).arr[i].index], (*ih).hash[(*ih).arr[j].index] = (*ih).hash[(*ih).arr[j].index], (*ih).hash[(*ih).arr[i].index]
	(*ih).arr[i], (*ih).arr[j] = (*ih).arr[j], (*ih).arr[i]
}

func (ih *intHeap) Pop() interface{} {
	ele := (*ih).arr[(*ih).Len() - 1]
	delete(ih.hash, (*ih).arr[(*ih).Len() - 1].index)
	(*ih).arr = (*ih).arr[0: (*ih).Len() - 1]
	return ele
}

func (ih *intHeap) Front() interface{} {
	ele := heap.Pop(ih)
	heap.Push(ih, ele)
	return ele
}

func (ih *intHeap) Remove(index int) {
	if value, exist := ih.hash[index]; exist {
		heap.Remove(ih, value)
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	results := make([]float64, 0)
	if len(nums) == 0 || k <= 0 || len(nums) < k {
		return []float64{}
	} else if k == 1 {
		for _, num := range nums {
			results = append(results, float64(num))
		}
		return results
	}

	smallHeap := &intHeap{
		arr: make([]*Element, 0),
		hash: map[int]int{},
		isSmall: true,
	}
	bigHeap := &intHeap{
		arr: make([]*Element, 0),
		hash: map[int]int{},
		isSmall: false,
	}
	for i, num := range nums {
		if smallHeap.Len() == 0 && bigHeap.Len() == 0 {
			heap.Push(bigHeap, &Element{
				num,
				i,
			})
		} else if smallHeap.Len() == 0 {
			if bigHeap.Front().(*Element).value > num {
				heap.Push(bigHeap, &Element{
					num,
					i,
				})
			} else {
				heap.Push(smallHeap, &Element{
					num,
					i,
				})
			}
		} else if bigHeap.Len() == 0 {
			if smallHeap.Front().(*Element).value > num {
				heap.Push(bigHeap, &Element{
					num,
					i,
				})
			} else {
				heap.Push(smallHeap, &Element{
					num,
					i,
				})
			}
		} else {
			if smallHeap.Front().(*Element).value >= num {
				heap.Push(bigHeap, &Element{
					num,
					i,
				})
			} else {
				heap.Push(smallHeap, &Element{
					num,
					i,
				})
			}
		}
		if i >= k {
			smallHeap.Remove(i-k)
			bigHeap.Remove(i-k)
		}
		balance(smallHeap, bigHeap)
		if i >= k - 1 {
			if smallHeap.Len() == bigHeap.Len() {
				results = append(results, float64(smallHeap.Front().(*Element).value + bigHeap.Front().(*Element).value) / 2.0 )
			} else {
				if smallHeap.Len() > bigHeap.Len() {
					results = append(results, float64(smallHeap.Front().(*Element).value))
				} else {
					results = append(results, float64(bigHeap.Front().(*Element).value))
				}
			}
		}
	}
	return results
}

// balance between small and big heap
func balance(small *intHeap, big *intHeap) {
	if small.Len() > big.Len() {
		balance(big, small)
	} else if small.Len() == big.Len() || small.Len() == big.Len() - 1 {
		return
	} else {
		for big.Len() - small.Len() > 1 {
			heap.Push(small, heap.Pop(big))
		}
	}
}
