package main

import "fmt"

func main() {
	nums := []int{3, 4, 9, 9, 4}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {

	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
	}
	for key, num := range set {
		if num == 1 {
			return key
		}
	}
	return 0
}
