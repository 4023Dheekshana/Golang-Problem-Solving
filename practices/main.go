package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []int{8, 3, 5, 5, 3, 8}
	arr2 := []int{3, 6, 4, 5, 2, 8}
	megedArray := mergeUnique(arr, arr2)
	fmt.Println(megedArray)
	a1 := []int{1, 2, 3}
	a2 := []int{4, 5, 6}
	dupes, uniques := duplicates(arr)
	fmt.Println("The uniques values are", uniques)
	fmt.Println("The dupes values are", dupes)
	merged := mergeArr(a1, a2)
	fmt.Println("The merged array is ", merged)
	str := "madam"
	fmt.Println(palindromecheck(str))
	list := []int{4, 2, 3, 1, 5}
	fmt.Println(missingNumber(list, 6))
	fmt.Println(conVow(str))
}
func duplicates(arr []int) ([]int, []int) {
	duplicates := []int{}
	uniques := []int{}
	countValues := make(map[int]int)
	for _, val := range arr {
		countValues[val]++
		if countValues[val] == 1 {
			uniques = append(uniques, val)
		} else if countValues[val] >= 2 {
			duplicates = append(duplicates, val)
		}
	}
	return duplicates, uniques
}

// merge two sorted arrays in a single single
func mergeArr(a1, a2 []int) []int {
	result := make([]int, 0, len(a1)+len(a2))
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			result = append(result, a1[i])
			i++
		} else {
			result = append(result, a2[j])
			j++
		}
	}
	result = append(result, a1[i:]...)
	result = append(result, a2[j:]...)

	return result
}

//merge two unsorted array while removing duplicates

func mergeUnique(arr1, arr2 []int) []int {
	combined := append(arr1, arr2...)
	sort.Ints(combined)
	fmt.Println(combined)
	result := []int{}
	for i, val := range combined {
		if i == 0 || val != combined[i-1] {
			result = append(result, val)
		}
	}
	return result
}

func palindromecheck(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func missingNumber(arr []int, n int) int {
	total := n * (n + 1) / 2

	expectedSum := 0
	for _, val := range arr {
		expectedSum += val
	}

	missingNum := total - expectedSum
	return missingNum

}

func conVow(lines string) (int, int) {
	vowels := "aeiouAEIOU"
	vowelCount := 0
	consonantCount := 0

	for _, char := range lines {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if char >= 'a' && char <= 'z' || char <= 'A' && char <= 'Z' {
			consonantCount++
		}
	}
	return consonantCount, vowelCount
}
