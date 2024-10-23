package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 5
	for i := 1; i <= n; i++ {
		spaces := strings.Repeat(" ", n-i)
		stars := strings.Repeat("*", 2*i-1)
		fmt.Println(spaces + stars)
	}
}
