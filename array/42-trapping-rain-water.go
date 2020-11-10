package main

import (
	"container/heap"
	"container/list"
)

// 42 trap rain water I: using two pointers
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	maxLeft, maxRight := 0, 0
	i, j := 0, len(height) - 1
	for i < j {
		if height[i] < height[j] {
			if height[i] > maxLeft {
				maxLeft = height[i]
			} else {
				res += maxLeft - height[i]
			}
			i += 1
		} else {
			if height[j] > maxRight {
				maxRight = height[j]
			} else {
				res += maxRight - height[j]
			}
			j -= 1
		}
	}
	return res
}

// 407 trapping rain water

// represent a point in height map
type Point struct {
	height int
	xPos   int
	yPos   int
}

// points priority queue
type Points []*Point

func(p *Points) Push(ele interface{}) {
	*p = append(*p, ele.(*Point))
}

func(p *Points) Pop() interface{}{
	ele := (*p)[len(*p)-1]
	*p = (*p)[0: len(*p)-1]
	return ele
}

func(p *Points) Less(i, j int) bool {
	return (*p)[i].height < (*p)[j].height
}

func(p *Points) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func(p *Points) Len() int {
	return len(*p)
}

// trap rain water II using heap and queue
func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	points := new(Points)
	heap.Init(points)
	visited := make([][]bool, 0)

	for i := 0; i < len(heightMap); i ++ {
		visited = append(visited, make([]bool, len(heightMap[0]), len(heightMap[0])))
	}
	// put all edge nodes in heap
	for i := 0; i < len(heightMap); i ++ {
		for j := 0; j < len(heightMap[0]); j ++ {
			if i == 0 || j == 0 || i == len(heightMap) - 1 || j == len(heightMap[0]) - 1 {
				heap.Push(points, &Point{
					heightMap[i][j],
					i,
					j,
				})
				visited[i][j] = true
			}
		}
	}

	res := 0

	for points.Len() > 0 {
		queue := list.New()
		low := heap.Pop(points).(*Point)
		minHeight := low.height
		for _, pos := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if isValid(low.xPos+pos[0], low.yPos+pos[1], len(heightMap), len(heightMap[0])) {
				if exist := visited[low.xPos+pos[0]][low.yPos+pos[1]]; !exist {
					visited[low.xPos+pos[0]][low.yPos+pos[1]] = true
					if heightMap[low.xPos+pos[0]][low.yPos+pos[1]] < minHeight {
						res += minHeight - heightMap[low.xPos+pos[0]][low.yPos+pos[1]]
						queue.PushBack(&Point{
							heightMap[low.xPos+pos[0]][low.yPos+pos[1]],
							low.xPos + pos[0],
							low.yPos + pos[1],
						})
					} else {
						heap.Push(points, &Point{
							heightMap[low.xPos+pos[0]][low.yPos+pos[1]],
							low.xPos + pos[0],
							low.yPos + pos[1],
						})
					}
				}
			}
		}

		for queue.Len() > 0 {
			low := queue.Front().Value.(*Point)
			queue.Remove(queue.Front())
			for _, pos := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				if isValid(low.xPos+pos[0], low.yPos+pos[1], len(heightMap), len(heightMap[0])) {
					if exist := visited[low.xPos+pos[0]][low.yPos+pos[1]]; !exist {
						visited[low.xPos+pos[0]][low.yPos+pos[1]] = true
						if heightMap[low.xPos+pos[0]][low.yPos+pos[1]] < minHeight {
							res += minHeight - heightMap[low.xPos+pos[0]][low.yPos+pos[1]]
							queue.PushBack(&Point{
								heightMap[low.xPos+pos[0]][low.yPos+pos[1]],
								low.xPos + pos[0],
								low.yPos + pos[1],
							})
						} else {
							heap.Push(points, &Point{
								heightMap[low.xPos+pos[0]][low.yPos+pos[1]],
								low.xPos + pos[0],
								low.yPos + pos[1],
							})
						}
					}
				}
			}
		}
	}
	return res
}

func isValid(x, y, length, width int) bool {
	return x >= 0 && y >= 0 && x < length && y < width
}
