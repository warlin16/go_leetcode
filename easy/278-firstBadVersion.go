package easy

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return true
}

// func firstBadVersion(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	bv := 0
// 	for i := n; i > 0; i-- {
// 		if isBadVersion(i) == false {
// 			bv = i
// 			return bv
// 		}
// 	}
// 	return bv
// }

// ^ brute force

func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
