package main

import (
	"fmt"
	"os"
	"strconv"
)

//declare the stack a
type Stack struct {
	elements []int
}

//
func (s *Stack) push(value int) {
	s.elements = append([]int{value}, s.elements...)
}

func (s *Stack) pop() int {
	if len(s.elements) == 0 {
		return 0
	}
	value := s.elements[0]
	s.elements = s.elements[1:]
	return value
}

func (s *Stack) top() int {
	if len(s.elements) == 0 {
		return 0
	}
	return s.elements[0]
}

func (s *Stack) swap() {
	if len(s.elements) < 2 {
		return
	}
	s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
}

func (s *Stack) rotate() {
	if len(s.elements) < 2 {
		return
	}
	first := s.elements[0]
	s.elements = append(s.elements[1:], first)
}

func (s *Stack) reverseRotate() {
	if len(s.elements) < 2 {
		return
	}
	last := s.elements[len(s.elements)-1]
	s.elements = append([]int{last}, s.elements[:len(s.elements)-1]...)
}

func sortStack(a, b *Stack) []string {
	var instructions []string
	// Simple sorting algorithm for demonstration purposes
	// This part should be replaced with a more optimized sorting algorithm
	for len(a.elements) > 0 {
		b.push(a.pop())
		instructions = append(instructions, "pb")
	}
	for len(b.elements) > 0 {
		a.push(b.pop())
		instructions = append(instructions, "pa")
	}
	return instructions
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	a := &Stack{}
	b := &Stack{}

	args := os.Args[1:]
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.push(num)
	}

	instructions := sortStack(a, b)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}

