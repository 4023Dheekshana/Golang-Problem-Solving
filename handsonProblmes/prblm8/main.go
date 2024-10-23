package main

//Group Anagrams
//Given an array of strings, group the anagrams together.
//An anagram is a word or phrase formed by rearranging the letters of another,
//such as cinema and iceman. Return the groups as an array of string slices.

import (
	"fmt"
	"sort"
)

func main() {
	anaInputs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(anaInputs))

}

func groupAnagrams(input []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range input {
		sortedStr := sortStr(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	result := [][]string{}
	for _, val := range anagrams {
		result = append(result, val)
	}
	return result

}

func sortStr(str string) string {
	sorted := []rune(str)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return string(sorted)
}
