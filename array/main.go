package main

import "fmt"

func main() {
	// dont delcare the size in the bracker otherwise it will become array
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	var nums =make([] int, 3,5)
	fmt.Println(cap(nums))

}