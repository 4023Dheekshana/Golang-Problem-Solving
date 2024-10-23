package main

import (
	"fmt"
	"math"
)

func largestOddNumber(num string) string {

	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return num[:i+1]
		}
	}
	return ""
}

func powerOfTwo(n int) bool {
	return math.Pow(float64(n), float64(2)) == float64(n)

}
func main() {
	fmt.Println(largestOddNumber("34"))
	fmt.Println(powerOfTwo(16))
}
