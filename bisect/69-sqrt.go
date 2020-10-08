package main

// last position of target * target <= x
func mySqrt(x int) int {
	start := 1
	end := x
	for start + 1 < end {
		mid := start + (end - start) / 2
		if mid * mid == x {
			return mid
		} else if mid * mid < x {
			start = mid
		} else {
			end = mid
		}
	}
	if end * end <= x {
		return end
	}
	return start
}


// follow up: double mySqrt
func myDoubleSqrt(x int) float64 {
	var start, end, mid float64
	start = 1.0
	end = float64(x)
	diff := 1e-06
	for start < end - diff {
		mid = start + (end - start) / 2
		if mid * mid == float64(x) {
			return mid
		} else if mid * mid < float64(x) {
			start = mid
		} else {
			end = mid
		}
	}
	if end * end <= float64(x) {
		return end
	}
	return start
}
