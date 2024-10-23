package main

import "fmt"

//rearrange array by alternating positive and negative numbers consecutively.

func rearrangeArray(nums []int) []int {
	pos := []int{}
	neg := []int{}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			neg = append(neg, nums[i])
		} else {
			pos = append(pos, nums[i])
		}
	}
	for p := 0; p < len(pos); {
		for n := 0; n < len(neg); {
			result = append(result, pos[p])
			p++
			result = append(result, neg[n])
			n++
		}
	}
	return result
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}
