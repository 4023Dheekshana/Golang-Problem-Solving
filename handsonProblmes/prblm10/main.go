package main

import "fmt"

// Product of Array Except Self
//Given an array nums of integers, return an array answer
//such that answer[i] is equal to the product of all the elements of nums except nums[i].
func main() {
	nums := []int{1, 5, 6, 3}
	fmt.Println(productByItself(nums))
}

func productByItself(nums []int) []int {

	n := len(nums)
	answer := make([]int, n)

	// Initialize answer array with left products
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// Compute right products and multiply with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer

}
