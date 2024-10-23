package main

import "fmt"

// Write a function in Go that takes a string and
// returns the first non-repeated (non-duplicate) character in it.
// If all characters are repeated or the string is empty, return a space character ' '.

func main() {
	ch := uniqcharacter("swiss")
	fmt.Println("The first non-repeated character is", ch)
}

func uniqcharacter(str string) string {
	uniqstring := make(map[rune]int)
	for _, ch := range str {
		uniqstring[ch]++
	}

	for key, val := range uniqstring {
		if val == 1 {
			return string(key)
		}
	}
	return ""
}
