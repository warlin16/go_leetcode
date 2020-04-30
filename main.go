package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // -> 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // -> 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // -> 3
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // -> 3
}
