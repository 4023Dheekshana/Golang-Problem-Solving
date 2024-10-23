package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func myAtoi(str string) int {
	s := strings.Trim(str, " ")
	var start int
	signMultiplier := 1
	n := len(s)
	if n == 0 {
		return 0
	}

	if s[0] == '-' {
		signMultiplier = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	var result int
	for i := start; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return result * signMultiplier
		}

		result = result*10 + int(s[i]-'0')

		if result*signMultiplier <= math.MinInt32 {
			return math.MinInt32
		} else if result*signMultiplier >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return result * signMultiplier
}
func main() {
	fmt.Println(greeting())
	fmt.Println(myAtoi("  -012"))
	CheckArm(100)
}

func greeting() string {
	g := []string{"Hello", "Hi", "Howdy"}
	return g[rand.Intn(len(g)-1)]
}

func CheckArm(number int) {
	reminder := 0
	temp := 0
	result := 0
	temp = number

	for {
		reminder = temp % 10
		result += reminder * reminder * reminder
		temp = temp / 10
		if temp == 0 {
			break
		}
	}
	if result == number {
		fmt.Println("The number is armstrong number")
	} else {
		fmt.Println("The number is not an armstrong number")
	}
}
