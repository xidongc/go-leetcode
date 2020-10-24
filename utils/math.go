package utils

func Max(a... int) int {
	size := len(a)
	if size == 0 {
		panic("input size equals to zero")// input illegal
	}
	maxVal := a[0]
	for i := 1; i < size; i ++ {
		if a[i] > maxVal  {
			maxVal = a[i]
		}
	}
	return maxVal
}

func Min(a... int) int {
	size := len(a)
	if size == 0 {
		panic("input size equals to zero") // input illegal
	}
	minVal := a[0]
	for i := 1; i < size; i ++ {
		if a[i] < minVal  {
			minVal = a[i]
		}
	}
	return minVal
}
