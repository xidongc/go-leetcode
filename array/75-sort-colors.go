package main

// 75 sort colors
func sortColors(nums []int)  {
	i, j := 0, len(nums) - 1
	for i < j {
		if nums[i] == 0 {
			i ++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j --
		}
	}
	for i < len(nums) && nums[i] == 0 {
		i += 1
	}
	j = len(nums) - 1
	for i < j {
		if nums[i] == 1 {
			i ++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j --
		}
	}
}

func main() {
	input := []int{2,0,1}
	sortColors(input)
}