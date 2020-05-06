package medium

// [3, 2, 3]
// [1, 2, 1, 2, 3, 1, 2, 1]

func majorityElement(nums []int) []int {
	count, count2, result, result2 := 0, 0, 0, 0
	for _, num := range nums {
		if num == result {
			count++
		} else if num == result2 {
			count2++
		} else if count == 0 {
			result = num
			count = 1
		} else if count2 == 0 {
			result2 = num
			count2 = 1
		} else {
			count--
			count2--
		}
	}

	output := []int{}
	sliceLen := len(nums) / 3
	count, count2 = 0, 0
	for _, num := range nums {
		if num == result {
			count++
		} else if num == result2 {
			count2++
		}
	}
	if count > sliceLen {
		output = append(output, result)
	}
	if count2 > sliceLen {
		output = append(output, result2)
	}
	return output
}
