package easy

func numJewelsInStones(j string, s string) int {
	jewels := map[string]int{}
	numOfJewels := 0
	for _, v := range s {
		jewels[string(v)]++
	}
	for _, v := range j {
		numOfJewels += jewels[string(v)]
	}
	return numOfJewels
}
