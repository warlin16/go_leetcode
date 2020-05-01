package main

/**
Given a string,
find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

// func lengthOfLongestSubstring(s string) int {
// 	m := make(map[string]int)
// 	increment := 0
// 	longest := 0
// 	for i := 0; i < len(s); i++ {
// 		letter := string(s[i])
// 		if m[letter] != 0 {
// 			i = increment
// 			m = make(map[string]int)
// 			increment++
// 			continue
// 		}
// 		m[letter]++
// 		if len(m) > longest {
// 			longest = len(m)
// 		}
// 	}
// 	return longest
// }

// brute force

func lengthOfLongestSubstring(s string) int {
	var current, longest int
	var m = make(map[rune]int)
	for i, char := range s {
		if start1, found := m[char]; found {
			if start1+1 > current {
				current = start1 + 1
			}
		}
		if longest < i-current+1 {
			longest++
		}
		m[char] = i
	}
	return longest
}
