package easy

func reverseString(s []byte) []byte {
	for i := 0; i < len(s)/2; i++ {
		leftAsc := s[i]
		rightAsc := s[len(s)-1-i]
		s[i], s[len(s)-1-i] = rightAsc, leftAsc
	}
	return s
}
