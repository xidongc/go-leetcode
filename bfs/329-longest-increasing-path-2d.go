package main

import (
	"container/heap"
	"container/list"
)

// 329 longest increasing path in a matrix using heap && queue (memorized bfs)
// the solution is similar with array/trapping-rain-water II (follow up), using
// heap to get mini value, and start from mini value, do bfs traversal
type Point struct {
	value int
	posX  int
	posY  int
}

type Points []*Point

func (p *Points) Len() int {
	return len(*p)
}

func (p *Points) Swap(a, b int) {
	(*p)[a], (*p)[b] = (*p)[b], (*p)[a]
}

func (p *Points) Push(ele interface{}) {
	*p = append(*p, ele.(*Point))
}

func (p *Points) Pop() interface{} {
	ele := (*p)[(*p).Len() - 1]
	*p = (*p)[:(*p).Len() - 1]
	return ele
}

func (p *Points) Less(a, b int) bool {
	return (*p)[a].value < (*p)[b].value
}
// dp[i][j] = max(dp[i-1][j]
func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, 0)
	for i := 0; i < len(matrix); i ++ {
		dp = append(dp, make([]int, len(matrix[0]), len(matrix[0])))
	}
	h := new(Points)
	for i := 0; i < len(matrix); i ++ {
		for j := 0; j < len(matrix[0]); j ++ {
			dp[i][j] = 1
			heap.Push(h, &Point{matrix[i][j], i, j})
		}
	}
	for h.Len() > 0 {
		ele := heap.Pop(h).(*Point)
		if dp[ele.posX][ele.posY] != 1 {
			continue
		}
		queue := list.New()
		queue.PushBack(ele)
		for queue.Len() > 0 {
			curr := queue.Front().Value.(*Point)
			queue.Remove(queue.Front())
			for _, pos := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				if valid(curr.posX+pos[0], curr.posY+pos[1], len(matrix), len(matrix[0])) &&
					matrix[curr.posX+pos[0]][curr.posY+pos[1]] > matrix[curr.posX][curr.posY] &&
					dp[curr.posX+pos[0]][curr.posY+pos[1]] < dp[curr.posX][curr.posY] + 1 {
					dp[curr.posX+pos[0]][curr.posY+pos[1]] = dp[curr.posX][curr.posY] + 1
					queue.PushBack(&Point{
						value: matrix[curr.posX+pos[0]][curr.posY+pos[1]],
						posX:  curr.posX+pos[0],
						posY:  curr.posY+pos[1],
					})
				}
			}
		}
	}

	maxLength := 0
	for i := 0; i < len(dp); i ++ {
		for j := 0; j < len(dp[0]); j ++ {
			if dp[i][j] > maxLength {
				maxLength = dp[i][j]
			}
		}
	}
	return maxLength
}

// if position valid
func valid(posX, posY, length, width int) bool {
	return posX >= 0 && posY >= 0 && posX < length && posY < width
}
