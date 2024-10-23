package main

import "fmt"

func main() {
	NewCalculation()
	n := 7
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func NewCalculation() {
	amt := 26000
	rent := 6000
	juicer := 2000
	myAmt := 7500
	AAmt := 5000
	balance := amt - (rent + juicer + myAmt + AAmt)
	fmt.Println(balance)
}
