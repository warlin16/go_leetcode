package easy

func numJewelsInStones(j string, s string) int {
	jewels := map[rune]int{}
	numOfJewels := 0
	for _, v := range s {
		jewels[v]++
	}
	for _, v := range j {
		numOfJewels += jewels[v]
	}
	return numOfJewels
}
