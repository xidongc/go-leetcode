package main

import (
	"container/list"
)

// 815 create bus -> bus solution
func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}
	graph := make(map[int][]int, 0)
	startSet := make([]int, 0)
	endSet := make([]int, 0)
	for i := 0; i < len(routes); i ++ {
		// determine whether in same bus, if so, return 1
		containsStart := contains(routes[i], S)
		containsEnd := contains(routes[i], T)
		if containsStart {
			startSet = append(startSet, i)
			if containsEnd {
				return 1
			}
		} else if containsEnd {
			endSet = append(endSet, i)
		}
		for j := i+1; j < len(routes); j ++ {
			if intersect(routes[i], routes[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	numBuses := 1

	// bfs
	queue := list.New()
	dedup := make(map[int]struct{}, 0)
	for _, startBus := range startSet {
		queue.PushBack(startBus)
		dedup[startBus] = struct{}{}
	}

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			ele := queue.Front().Value.(int)
			queue.Remove(queue.Front())
			if nextBuses, exist := graph[ele]; exist {
				for _, nextBus := range nextBuses {
					if contains(endSet, nextBus) {
						return numBuses + 1
					}
					if _, exist := dedup[nextBus]; !exist {
						queue.PushBack(nextBus)
						dedup[nextBus] = struct{}{}
					}
				}
			}
		}
		numBuses += 1
	}
	return -1
}

// Determine whether route1 and route2 has intersect using hashset O(n)
func intersect(route1 []int, route2 []int) bool {
	hashSet := make(map[int]struct{}, 0)
	for _, bus := range route1 {
		hashSet[bus] = struct{}{}
	}
	for _, bus := range route2 {
		if _, exist := hashSet[bus]; exist {
			return true
		}
	}
	return false
}

// return whether stop in route
func contains(route []int, stop int) bool {
	for _, bus := range route {
		if bus == stop {
			return true
		}
	}
	return false
}