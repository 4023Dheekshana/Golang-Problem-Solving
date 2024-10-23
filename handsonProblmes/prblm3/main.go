package main

import (
	"fmt"
)

type employee struct {
	name     string
	date     int
	location string
}

// type myInt = int

func main() {
	nrmlsplits()
	emp1 := employee{name: "rahul", date: 4, location: "chennai"}
	fmt.Println(emp1)

	// map1 := map[string]int{
	// 	"name":     10,
	// 	"location": 20,
	// }
	// fmt.Println(map1["name"])
	// // var i interface{} = 3
	// // f := i.(myInt)
	// // fmt.Println(f)

}

// func fizzbuzz() {
// 	for i := 1; i <= 10; i++ {
// 		if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("Fizz Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

func nrmlsplits() {
	mnthlySalary := 26000
	pgRent := 6200
	mySavings := 7500
	//debt := 5000
	//Aamt := 5000
	balance := mnthlySalary - (pgRent + mySavings)
	fmt.Println("Usual balance: ", balance)
}
