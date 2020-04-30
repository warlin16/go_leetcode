package main

/**
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed.
All houses at this place are arranged in a circle.
That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have security system connected and
it will automatically contact the police if two adjacent houses
were broken into on the same night.

Given a list of non-negative integers representing the amount of money
of each house, determine the maximum amount of money you can rob tonight
without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
because they are adjacent houses.
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
*/

// [0] -> 0
// [1, 2] -> 2
// [2, 1, 1, 2, 6] -> 7
// [3, 4, 3] -> 4
// [1, 5, 1] -> 5
// [1, 2, 1, 0] -> 2
// [1, 2, 3, 1] -> 4
// [2, 1, 1, 2] -> 3

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
