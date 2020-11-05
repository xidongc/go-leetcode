package utils

import (
	"container/heap"
	"strconv"
)

type Element struct {
	Index int
	Value int
}

func (ele *Element) Hash() string {
	return "s" + strconv.Itoa(ele.Index) + "ss" + strconv.Itoa(ele.Value)
}

type TreeMap struct {
	heap []*Element
	hash map[string]int
}

func InitTreeMap() *TreeMap{
	treeMap:= &TreeMap{
		heap: []*Element{},
		hash: map[string]int{},
	}
	heap.Init(treeMap)
	return treeMap
}

func (s *TreeMap) Front() *Element {
	ele := heap.Pop(s)
	heap.Push(s, ele)
	return ele.(*Element)
}

func (s *TreeMap) Len() int {
	return len(s.heap)
}

func (s *TreeMap) Swap(i, j int) {
	tmp := s.hash[s.heap[i].Hash()]
	s.hash[s.heap[i].Hash()] = s.hash[s.heap[j].Hash()]
	s.hash[s.heap[j].Hash()] = tmp
	s.heap[i], s.heap[j] = s.heap[j], s.heap[i]
}

func (s *TreeMap) Less(i, j int) bool {
	if s.heap[i].Value < s.heap[j].Value {
		return true
	}
	return false
}

func (s *TreeMap) Push(ele interface{}) {
	(*s).heap = append((*s).heap, ele.(*Element))
	(*s).hash[ele.(*Element).Hash()] = s.Len()-1
}

func (s *TreeMap) Pop() interface{} {
	ele := (*s).heap[len((*s).heap)-1]
	delete(s.hash, (*s).heap[len((*s).heap)-1].Hash())
	(*s).heap = (*s).heap[0:len((*s).heap)-1]
	return ele
}

func (s *TreeMap) Remove(ele *Element) {
	if idx, exist := s.hash[ele.Hash()]; exist {
		heap.Remove(s, idx)
	}
}
