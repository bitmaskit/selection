package selection

func Sort(nums []int) {
	for i := range nums {
		min := i + minIdx(nums[i:])
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func minIdx(nums []int) (min int) {
	for i := range nums {
		if nums[i] < nums[min] {
			min = i
		}
	}
	return
}
