package easy

func TwoSum(nums []int, target int) []int {
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
