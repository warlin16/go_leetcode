package main

import "fmt"

func main() {
	fmt.Println(":eyes:", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(":eyes:", containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(":eyes:", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 1, 4, 2}))
}
