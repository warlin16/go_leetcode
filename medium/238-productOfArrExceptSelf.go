package medium

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
