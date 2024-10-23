package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 6}
	target := 3
	fmt.Println(binarySearch(arr, target))
	fmt.Println(linearSearch(arr, target))
}

func linearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		median := (low + high) / 2
		if arr[median] == target {
			return median
		}
		if arr[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	return -1

}
