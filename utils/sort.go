package utils

func RadixSort(nums *[]int) {
	shrink := 1
	maxValue := Max(*nums...)

	for maxValue > 0 {
		var tmp = [10][]int{}
		for _, num := range *nums {
			original := num
			num = num / shrink
			tmp[num%10] = append(tmp[num%10], original)
		}
		shrink *= 10

		i := 0
		for i < len(*nums) {
			for _, list := range tmp {
				if len(list) > 0 {
					for _, num := range list {
						(*nums)[i] = num
						i += 1
					}
				}
			}
		}
		maxValue /= 10
	}
}
