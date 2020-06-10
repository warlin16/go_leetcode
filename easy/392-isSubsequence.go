package easy

func isSubsequence(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	si := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[si] {
			si++
		}
		if si == len(s) {
			return true
		}
	}
	return false
}
