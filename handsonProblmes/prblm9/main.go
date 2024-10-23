package main

import "fmt"

//Find the Longest Substring Without Repeating Characters
//Given a string, find the length of the longest substring without repeating characters.

//Window sliding approach which has left and right pointer
//window sliding can be used in the list, array, and maps.

func main() {
	inputString := "abcdefgwlabaab"
	fmt.Println(LengthOfLongestSubstring(inputString))
}

func LengthOfLongestSubstring(str string) int {
	charIndex := make(map[rune]int)
	left := 0
	maxLength := 0

	for right, char := range str {
		if index, found := charIndex[char]; found && index >= left {
			left = index + 1
		}
		charIndex[char] = right
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
