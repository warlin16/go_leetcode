package main

import "fmt"

/**
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/

var cache map[int]int = make(map[int]int)

func climbStairs(n int) int {
	fmt.Println("What is this value", cache[12])
	if n <= 1 {
		return 1
	}
	if cache[n] == 0 {
		cache[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return cache[n]
}

// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	numberOfSteps := []int{1, 2}
// 	for len(numberOfSteps) != n {
// 		lastNum := numberOfSteps[len(numberOfSteps)-1]
// 		penultimateNum := numberOfSteps[len(numberOfSteps)-2]
// 		numberOfSteps = append(numberOfSteps, lastNum+penultimateNum)
// 	}
// 	return numberOfSteps[len(numberOfSteps)-1]
// }
