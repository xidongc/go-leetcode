package main

import (
	"sort"
)

// 56 merge intervals: using line sweep algorithm
type TimePoint struct {
	time int
	isEnd bool
}

func (t *IntervalArr) Less(i, j int) bool {
	if (*t)[i].time < (*t)[j].time {
		return true
	} else if (*t)[i].time == (*t)[j].time {
		return (*t)[i].isEnd == false
	}
	return false
}

func (t *IntervalArr) Len() int {
 	return len(*t)
}

func (t *IntervalArr) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

type IntervalArr []*TimePoint

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	intervalsArr := make(IntervalArr, 0)
	for _, interval := range intervals {
		if len(interval) != 2 {
			panic("input ele size not equal to 2")
		}
		intervalsArr = append(intervalsArr, &TimePoint {
			interval[0],
			false,
		})
		intervalsArr = append(intervalsArr, &TimePoint {
			interval[1],
			true,
		})
	}

	sort.Sort(&intervalsArr)

	count := 1
	start := intervalsArr[0].time
	reachEnd := false
	results := make([][]int, 0)
	for i := 1; i < intervalsArr.Len(); i ++ {
		if reachEnd {
			start = intervalsArr[i].time
			reachEnd = false
		}
		if intervalsArr[i].isEnd == true {
			count -= 1
		} else {
			count += 1
		}
		if count == 0 {
			results = append(results, []int{start, intervalsArr[i].time})
			reachEnd = true
		}
	}
	return results
}
