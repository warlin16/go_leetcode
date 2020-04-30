package main

import "fmt"

func main() {
	fmt.Println(robTwo([]int{0}))
	fmt.Println(robTwo([]int{1, 2}))
	fmt.Println(robTwo([]int{2, 1, 1, 2, 6}))
	fmt.Println(robTwo([]int{3, 4, 3}))
	fmt.Println(robTwo([]int{1, 5, 1}))
	fmt.Println(robTwo([]int{1, 2, 1, 0}))
	fmt.Println(robTwo([]int{1, 2, 3, 1}))
	fmt.Println(robTwo([]int{2, 1, 1, 2}))
}
