package main

import "fmt"

func main() {
	arr := []int{3, 6, 1, 4, 4}
	fmt.Println(checkDuplicates(arr))
}

func checkDuplicates(arr []int) bool {
	mappedArr := make(map[int]bool)
	for _, num := range arr {
		if mappedArr[num] {
			return true
		}
		mappedArr[num] = true
	}
	return false
}
