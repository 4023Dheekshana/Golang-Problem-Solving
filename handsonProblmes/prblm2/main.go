// Write a function in Go that takes a string and
// returns the length of the longest substring without repeating characters.
package main

import (
	"fmt"
)

func main() {
	//fmt.Println(longestSubstring("abcabcbb"))
	majorityElement([]int{2, 6, 4, 7, 3, 3, 3})
	mapFeature()
	// x := []int{1, 2, 3, 4}
	// y := []int{5, 6, 7, 9}
	// fmt.Println(mergeSort(x, y, len(x), len(y)))
}

// func longestSubstring(str string) int {
// 	longSS := make(map[rune]bool)
// 	uniqSubstring := []string{}
// 	for _, ch := range str {
// 		longSS[ch] = true
// 		if longSS[ch] {
// 			uniqSubstring = append(uniqSubstring, string(ch))
// 		}
// 	}
// 	return len(uniqSubstring)
// }

func majorityElement(nums []int) {
	count := make(map[int]int)
	for _, val := range count {
		count[val]++
	}
	for key, val := range count {
		if val > len(nums)/2 {
			fmt.Println(key)
		}
	}

}

func mapFeature() {
	//map is unordered but while printing it prints in ascending order
	//because of go's internal implementation which organize the elements in order during runtime and
	//it is not assured that it will always print in the ascending order.
	varMap := map[string]string{
		"name":     "Dheeksh",
		"age":      "22",
		"location": "Chennai",
		"board":    "StateBoard",
	}
	fmt.Println(varMap)

}
