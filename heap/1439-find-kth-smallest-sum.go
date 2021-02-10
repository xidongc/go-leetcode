package main

import (
	"container/heap"
)

// 1439 find kth smallest sum of a matrix with sorted rows: using heap, and n-dimension matrix
type Elements struct {
	value int
	pos []int
}

type PriorityQueue []*Elements

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].value < (*p)[j].value
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*Elements))
}

func (p *PriorityQueue) Pop() interface{} {
	last := (*p)[len(*p) - 1]
	*p = (*p)[: len(*p) - 1]
	return last
}

// kth smallest
func kthSmallest(mat [][]int, k int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		panic("invalid input matrix")
	}
	sum := 0
	current := 0
	pq := new(PriorityQueue)
	pos := make([]int, 0)
	for i := 0; i < len(mat); i ++ {
		sum += mat[i][0]
		pos = append(pos, 0)
	}
	visited := map[int]struct{}{}
	heap.Push(pq, &Elements{
		value: sum,
		pos: pos,
	})
	visited[indexHash(pos)] = struct{}{}
	for current < k - 1 {
		ele := heap.Pop(pq).(*Elements)
		current += 1
		for i := 0; i < len(ele.pos); i ++ {
			tmp := make([]int, len(ele.pos), len(ele.pos))
			copy(tmp, ele.pos)
			tmp[i] += 1

			if _, exist := visited[indexHash(tmp)]; exist || tmp[i] >= len(mat[0]) {
				continue
			}
			heap.Push(pq, &Elements{
				value: getSum(tmp, mat),
				pos: tmp,
			})
			visited[indexHash(tmp)] = struct{}{}
		}
	}
	return heap.Pop(pq).(*Elements).value
}

// get sum
func getSum(index []int, matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix); i ++ {
		sum += matrix[i][index[i]]
	}
	return sum
}

func indexHash(index []int) int {
	sum := 0
	mul := 1
	for i := len(index) - 1; i >= 0; i -- {
		sum += index[i] * mul
		mul *= 10
	}
	return sum
}
