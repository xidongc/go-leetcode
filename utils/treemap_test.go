package utils

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestTreeMap(t *testing.T) {
	fmt.Println("start program... ")
	input := []*Element{{7, 0}, {7, 1}, {6, 2}, {5, 3}, {4, 4}, {3, 5}, {2, 6}, {1, 7}}
	treeMap := InitTreeMap()
	for _, ele := range input {
		heap.Push(treeMap, ele)
	}

	treeMap.Remove(&Element{Index: 3, Value: 5})

	for treeMap.Len() > 0 {
		frontEle := treeMap.Front()
		popEle := heap.Pop(treeMap)

		if frontEle != popEle {
			t.Error("ele value not match")
		}
	}
}
