package main

import "fmt"

type stack struct {
	slice []int
}

func (s *stack) push(val int) {
	s.slice = append(s.slice, val)
}

func (s *stack) pop() int {
	n := len(s.slice)
	if n == 0 {
		fmt.Println("The stack is empty")
	}
	poppedElement := s.slice[n-1]
	s.slice = s.slice[:n-1]
	return poppedElement
}
func (s *stack) peek() int {
	n := len(s.slice)
	peekElement := s.slice[n-1]
	return peekElement
}

func main() {
	sc := stack{}
	sc.push(9)
	sc.push(6)
	fmt.Println(sc.peek())
	fmt.Println(sc.pop())
}
