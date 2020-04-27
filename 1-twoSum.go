package main

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, AND YOU MAY NOT USE THE SAME ELEMENT TWICE.
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		m[num] = i
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if m[target-num] != i && m[target-num] != 0 {
			return []int{i, m[target-num]}
		}
	}
	return []int{}
}
