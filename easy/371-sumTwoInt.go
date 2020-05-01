package easy

func getSum(a int, b int) int {
	for b != 0 {
		c := a & b
		a ^= b
		b = c << 1
	}
	return a
}
