package main

import (
	"container/list"
)

// 210 course schedule II
// Note: init will avoid corner case where len(inDegree) < numCourses
func findOrderII(numCourses int, prerequisites [][]int) []int {

	sequence := make([]int, 0)

	if numCourses == 0 {
		return sequence
	}

	inDegree := make(map[int]int)
	queue := list.New()
	graph := make(map[int][]int)

	// init
	for i := 0; i <numCourses; i ++ {
		inDegree[i] = 0
		graph[i] = []int{}
	}

	// calculate in degree for each node
	for _, prerequisite := range prerequisites {
		if len(prerequisite) != 2 {
			panic("invalid input")
		}
		if _, exist := inDegree[prerequisite[1]]; exist {
			graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		}

		if degree, exist := inDegree[prerequisite[0]]; exist {
			inDegree[prerequisite[0]] = degree + 1
		}
	}

	// find in-degree == 0
	for node, degree := range inDegree {
		if degree == 0 {
			queue.PushBack(node)
		}
	}

	// bfs
	for queue.Len() > 0 {
		ele := queue.Front().Value.(int)
		sequence = append(sequence, ele)
		queue.Remove(queue.Front())
		for _, child := range graph[ele] {
			inDegree[child] -= 1
			if inDegree[child] == 0 {
				queue.PushBack(child)
			}
		}
	}

	// check if possible to finish all courses
	if len(sequence) == numCourses {
		return sequence
	}
	return []int{}
}
