package utils

type Ele struct {
	value 	 int
	priority int
}

type PriorityQueue []*Ele

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Ele))
}

