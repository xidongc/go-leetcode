package main

import (
	"container/list"
)

/*
210 course schedule II

There are a total of n courses you have to take labelled from 0 to n - 1.
Some courses may have prerequisites, for example, if prerequisites[i] = [ai, bi] this means you must take the course bi before the course ai.
Given the total number of courses numCourses and a list of the prerequisite pairs, return the ordering of courses you should take to finish all courses.
If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

topological sorting, find nodes where in-degree == 0

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:
Input: numCourses = 1, prerequisites = []
Output: [0]
*/
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
