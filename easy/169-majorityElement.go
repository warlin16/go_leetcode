package easy

// array of size n -> n / 2
// [3, 2, 3] -> n / 2 == 1.5 so majorityElement is 3 since it appears
// more than n / 2
// [2, 2, 1, 1, 1, 2, 2] -> 2

// func majorityElement(nums []int) int {
// 	numToCheck := len(nums) / 2
// 	m := make(map[int]int)
// 	for _, n := range nums {
// 		m[n]++
// 	}
// 	for k, v := range m {
// 		if v > numToCheck {
// 			return k
// 		}
// 	}
// 	return 0
// }

func majorityElement(nums []int) int {
	count := 0
	result := 0
	for _, n := range nums {
		if count == 0 {
			result = n
		}
		if n == result {
			count++
		} else {
			count--
		}
	}
	return result
}
