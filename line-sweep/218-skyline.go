package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
)

// 218 skyline, not passed due to some corner case, using hash heap with index
type Position struct {
	pos int
	height int
	isEnd bool
	index int
}

type SkylineHeap []*Position

func (s SkylineHeap) Len() int {
	return len(s)
}

func (s SkylineHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SkylineHeap) Less(i, j int) bool {
	if s[i].pos < s[j].pos {
		return true
	} else if s[i].pos == s[j].pos {
		if s[i].height > s[j].height {
			return true
		} else if s[i].height == s[j].height {
			return s[i].isEnd == false
		}
	}
	return false
}

type Skyline struct {
	heap []*Element
	hash map[string]int
}

func (s Skyline) Len() int {
	return len(s.heap)
}

func (s Skyline) Swap(i, j int) {
	tmp := s.hash[s.heap[i].Hash()]
	s.hash[s.heap[i].Hash()] = s.hash[s.heap[j].Hash()]
	s.hash[s.heap[j].Hash()] = tmp
	s.heap[i], s.heap[j] = s.heap[j], s.heap[i]
}

func (s Skyline) Less(i, j int) bool {
	if s.heap[i].height < s.heap[j].height {
		return true
	}
	return false
}

func (s *Skyline) Push(ele interface{}) {
	(*s).heap = append((*s).heap, ele.(*Element))
	(*s).hash[ele.(*Element).Hash()] = s.Len()-1
}

func (s *Skyline) Pop() interface{} {
	ele := (*s).heap[len((*s).heap)-1]
	delete(s.hash, (*s).heap[len((*s).heap)-1].Hash())
	(*s).heap = (*s).heap[0:len((*s).heap)-1]
	return ele
}

func (s *Skyline) Remove(ele *Element) {
	if idx, exist := s.hash[ele.Hash()]; exist {
		heap.Remove(s, idx)
	}
}

func getSkyline(buildings [][]int) [][]int {
	h := make(SkylineHeap, 0)
	for idx, building := range buildings {
		if len(building) != 3 {
			panic("building input invalid")
		}
		startPos := &Position{building[0], building[2], false, idx}
		endPos := &Position{building[1], building[2], true, idx}
		h = append(h, startPos)
		h = append(h, endPos)
	}
	sort.Sort(h)

	result := make([][]int, 0)
	highest := math.MinInt64

	skyline := Skyline{
		heap: []*Element{},
		hash: map[string]int{},
	}
	heap.Init(&skyline)

	for i := 0; i < len(h); i ++ {
		if h[i].isEnd == false {
			heap.Push(&skyline, &Element{h[i].index, h[i].height})
			if skyline.heap[skyline.Len()-1].height > highest {
				result = append(result, []int{h[i].pos, skyline.heap[skyline.Len()-1].height})
				highest = skyline.heap[skyline.Len()-1].height
			}
		} else {
			skyline.Remove(&Element{ h[i].index, h[i].height})
			if skyline.Len() > 0 {
				if skyline.heap[skyline.Len()-1].height < highest {
					result = append(result, []int{h[i].pos, skyline.heap[skyline.Len()-1].height})
					highest = skyline.heap[skyline.Len()-1].height
				}
			} else {
				result = append(result, []int{h[i].pos, 0})
				highest = 0
			}
		}
	}
	return result
}

type Element struct {
	index int
	height int
}

func (ele *Element) Hash() string {
	return "s" + strconv.Itoa(ele.index) + "ss" + strconv.Itoa(ele.height)
}

func main() {
	// [ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]
	// skyline2 := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	// skyline2 := [][]int{{0,2,3},{2,5,3}}
	skyline2 := [][]int{{2,4,7},{2,4,5},{2,4,6}}
	fmt.Println(getSkyline(skyline2))
}
