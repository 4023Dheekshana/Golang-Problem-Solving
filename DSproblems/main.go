package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println(productExceptSelf([]int{2, 4, 5, 6}))
	fmt.Println(matrix(4))
	fmt.Println(lengthOfLastWord("Hello world"))
	vowels := "Dheekshana"
	fmt.Println(reverseVowels(vowels))
	arr := []int{1, 2, 3, 1, 2, 2, 3}
	arrString := []string{"Dog", "Cat", "Lion", "Tiger"}
	fmt.Println(reverseAnArray(arrString))
	fmt.Println(reverseAnArray2(arr))
	fmt.Println(frequency(arr))
	str1 := "abs"
	str2 := "sba"
	palinString := "MAdam"
	palinNumber := "98709"
	fmt.Println(anagrams(str1, str2))
	fmt.Println(duplicates(arr))
	arr1 := []int{1, 2, 3}
	arr2 := []int{3, 8, 9}
	fmt.Println(intersectionArray(arr1, arr2))
	fmt.Println(mergeArr(arr1, arr2))
	str := "Golang is open source language"
	fmt.Println(stringReverse1(str))
	fmt.Println(stringReverse2(str))
	fmt.Println(primeORnot(11))
	numbers := 121
	fmt.Println(palindrome(numbers))
	fmt.Println(palindromeString(palinString))
	fmt.Println(palindromeString(palinNumber))
	fmt.Println(factorial(8))
	a := "12"
	b := "12"
	fmt.Println(strToInteger(a, b))
	a1 := []int{1, 3, 5, 7}
	a2 := []int{2, 4, 6, 8}
	a1s := a1[:]
	a2s := a2[:]
	fmt.Println(appendTwoArry(a1s, a2s))
	swapa := 89
	swapb := 76
	fmt.Println(swap(swapa, swapb))
	// MoveZ := []int{1, 0, 3, 4, 0, 2}
	// fmt.Println(movezeros(MoveZ))

}
func appendTwoArry(a1, a2 []int) []int {
	sort.Ints(a1)
	sort.Ints(a2)
	result := append(a1, a2...)
	sort.Ints(result)
	return result
}

func frequency(arr []int) map[int]int {
	freqmap := make(map[int]int)
	for _, val := range arr {
		freqmap[val]++
	}
	return freqmap
}

func anagrams(str1, str2 string) bool {
	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")
	sort.Strings(s1)
	sort.Strings(s2)
	return strings.Join(s1, "") == strings.Join(s2, "")
}

func duplicates(arr []int) ([]int, []int) {
	countVal := make(map[int]int)
	uniques := []int{}
	duplicates := []int{}
	for _, val := range arr {
		countVal[val]++
		if countVal[val] == 1 {
			uniques = append(uniques, val)
		} else if countVal[val] >= 2 {
			duplicates = append(duplicates, val)
		}
	}
	return uniques, duplicates
}

func intersectionArray(arr1, arr2 []int) []int {
	set := make(map[int]bool)
	result := []int{}
	for _, val := range arr1 {
		set[val] = true
	}
	for _, val := range arr2 {
		if set[val] {
			result = append(result, val)
		}
	}
	return result
}

func mergeArr(arr1, arr2 []int) []int {
	mergedArr := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}
	mergedArr = append(mergedArr, arr1[i:]...)
	mergedArr = append(mergedArr, arr2[j:]...)
	return mergedArr
}

func stringReverse1(str string) string {
	reversed := ""
	for _, ch := range str {
		reversed = string(ch) + reversed
	}
	return reversed
}

func stringReverse2(str string) string {

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func primeORnot(val int) bool {
	if val <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(val)))
	for i := 2; i <= sqrt; i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func palindromeString(s string) bool {
	removed := make([]rune, 0, len(s))
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsNumber(ch) {
			removed = append(removed, unicode.ToLower(ch))
		}
	}
	for i := 0; i < len(removed)/2; i++ {
		if removed[i] != removed[len(removed)-1-i] {
			return false
		}
	}
	return true
}

func palindrome(n int) bool {
	original := n
	reversed := 0
	for n > 0 {
		ld := n % 10
		reversed = (reversed * 10) + ld
		n = n / 10
	}
	return original == reversed
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func reverseAnArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func reverseAnArray2(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func strToInteger(a, b string) int {
	Inta, _ := strconv.Atoi(a)
	Intb, _ := strconv.Atoi(b)
	result := Inta * Intb
	return result
}

func swap(a, b int) (int, int) {
	return b, a
}

// func movezeros(arr []int) []int {
// 	for i := 0; i <= len(arr)-1; i++ {
// 		if arr[i] == 0 {
// 			arr[len(arr)-1] = 1
// 			arr[len(arr)]--
// 		}
// 	}
// 	return arr
// }

func reverseVowels(s string) string {

	vowels := "aeiouAEIOU"
	slice := []rune(s)

	for i, j := 0, len(slice)-1; i < j; {
		if !strings.ContainsRune(vowels, slice[i]) {
			i++
			continue
		}
		if !strings.ContainsRune(vowels, slice[j]) {
			j--
			continue
		}
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--

	}
	return string(slice)
}

// func productExceptSelf(nums []int) []int {
// 	result := make([]int, 0, len(nums))
// 	product := 0
// 	arrayIndex := 0
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := 1; j < len(nums)-1; j++ {
// 			product *= nums[j]
// 			result[arrayIndex] = product
// 			arrayIndex++
// 		}
// 	}
// 	return result
// }

func lengthOfLastWord(s string) int {
	trimmed := strings.Fields(s)
	final := trimmed[len(trimmed)-1]
	result := len(final)

	return result
}

func matrix(n int) int {
	// similar to fibonacci series matrix of 4 is matrix of 3and2, matrix of 3 is matrix of 2and1,
	// matrix of 2 is matrix of 1and0
	if n < 2 {
		return n
	}
	return matrix(n-1) + matrix(n-2) // formula of fibonacci sequence: f(n) = f(n-1) + f(n-2)
}
