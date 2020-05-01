package easy

/**
Given a non-empty array of integers,
every element appears twice except for one.Find that single one.

Note:

Your algorithm should have a linear runtime complexity.
Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

// func singleNumber(nums []int) int {
// 	m := map[int]int{}
// 	for i := range nums {
// 		m[nums[i]]++
// 	}
// 	for k, v := range m {
// 		if v == 1 {
// 			return k
// 		}
// 	}
// 	return 0
// }

func singleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
