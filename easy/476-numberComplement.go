package easy

func findComplement(num int) int {
	result := 0
	base := 1
	for num > 0 {
		if num%2 == 0 {
			result += base
		}

		num = num >> 1
		if num == 0 {
			break
		}
		base = base << 1
	}
	return result
}
