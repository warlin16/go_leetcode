package medium

func robTwo(nums []int) int {
	if len(nums) <= 3 {
		return max(nums...)
	}
	result1 := helperThief(nums[:len(nums)-1])
	result2 := helperThief(nums[1:])
	if result1 > result2 {
		return result1
	}
	return result2
}

func helperThief(nums []int) int {
	previous := 0
	current := 0
	for _, value := range nums {
		temp := current
		if previous+value > current {
			current = previous + value
		}
		previous = temp
	}
	return current
}

func max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
