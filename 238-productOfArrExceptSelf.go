package main

/**
Given an array nums of n integers where n > 1,
return an array output such that output[i] is equal
to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Constraint: It's guaranteed that the product
of the elements of any prefix or suffix of the array
(including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity?
(The output array does not count as extra space for the
purpose of space complexity analysis.)
*/

// func productExceptSelf(nums []int) []int {
// 	m := make(map[int][]int)
// 	output := make([]int, len(nums))
// 	for i := range nums {
// 		m[i] = append(m[i], nums[:i]...)
// 		m[i] = append(m[i], nums[i+1:]...)
// 	}
// 	for k, v := range m {
// 		product := 1
// 		for _, num := range v {
// 			product *= num
// 		}
// 		output[k] = product
// 	}

// 	return output
// }

// ^^^ super brute force, didn't even pass submission
// input -> [1, 2, 3, 4]

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	leftArr := make([]int, len(nums))
	rightArr := make([]int, len(nums))
	leftArr[0] = 1
	rightArr[len(nums)-1] = 1
	for i, j := 1, len(nums)-2; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		leftArr[i] = leftArr[i-1] * nums[i-1]
		rightArr[j] = rightArr[j+1] * nums[j+1]
	}
	for i := range output {
		output[i] = leftArr[i] * rightArr[i]
	}
	return output
}
