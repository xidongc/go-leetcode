package main

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
)

// 4 median of two sorted array
// implement find k ele in nums1 + nums2 array, by check k/2 ele in both nums1 and nums2
// remove [0 : k/2] k/2 ele if its kth ele is smaller
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}
	size := len(nums1) + len(nums2)
	if size % 2 == 1 {
		return float64(findKEle(nums1, nums2, size/2 + 1))
	} else {
		return float64(findKEle(nums1, nums2, size/2) + findKEle(nums1, nums2, size/2+1)) / 2.0
	}
}

func findKEle(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	} else if len(nums2) == 0 {
		return nums1[k-1]
	} else if k == 1 {
		return utils.Min(nums1[0], nums2[0])
	}

	if len(nums1) < len(nums2) {
		return findKEle(nums2, nums1, k)
	}

	if len(nums1) > len(nums2) {
		size := len(nums1) - len(nums2)
		for i := 0; i < size; i ++ {
			nums2 = append(nums2, math.MaxInt64)
		}
	}

	if len(nums1) >= k/2 {
		if nums1[k/2-1] < nums2[k/2-1] {
			return findKEle(nums1[k/2:], nums2, k - k/2)
		} else {
			return findKEle(nums1, nums2[k/2:], k - k/2)
		}
	}
	return -1
}