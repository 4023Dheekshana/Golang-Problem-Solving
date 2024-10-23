package main

import (
	"fmt"
	"math"
)

func buySellStock(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				sub := float64(arr[j] - arr[i])
				res = int(math.Max(sub, float64((res))))
			}
		}
	}
	return res
}
func main() {
	fmt.Println(buySellStock([]int{7, 1, 5, 3, 6, 4}))
}
