package main

import (
	"fmt"
)

// Stack represents our push_swap stack
type Stack struct {
	elements []int
}

// Push adds an element to the stack
func (s *Stack) Push(val int) {
	s.elements = append([]int{val}, s.elements...)
}

// Pop removes and returns the top element
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1 // Stack is empty
	}
	val := s.elements[0]
	s.elements = s.elements[1:]
	return val
}

// Swap the top two elements
func (s *Stack) Swap() {
	if len(s.elements) < 2 {
		return
	}
	s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
	fmt.Println("sa")
}

// Rotate the stack upwards
func (s *Stack) Rotate() {
	if len(s.elements) < 2 {
		return
	}
	first := s.elements[0]
	s.elements = append(s.elements[1:], first)
	fmt.Println("ra")
}

// Reverse rotate the stack downwards
func (s *Stack) ReverseRotate() {
	if len(s.elements) < 2 {
		return
	}
	last := s.elements[len(s.elements)-1]
	s.elements = append([]int{last}, s.elements[:len(s.elements)-1]...)
	fmt.Println("rra")
}

// Sorts stack of size 2
func sortTwo(s *Stack) {
	if s.elements[0] > s.elements[1] {
		s.Swap()
	}
}

// Sorts stack of size 3
func sortThree(s *Stack) {
	a, b, c := s.elements[0], s.elements[1], s.elements[2]
	if a > b && b > c { // 321
		s.Swap()
		s.ReverseRotate()
	} else if a > c && c > b { // 312
		s.Rotate()
	} else if b > a && a > c { // 231
		s.ReverseRotate()
	} else if b > c && c > a { // 213
		s.Swap()
	} else if c > a && a > b { // 132
		s.Swap()
		s.Rotate()
	}
}

// Sorts stack of size 4 or 5 using push to `b` strategy
func sortFourOrFive(s *Stack) {
	b := &Stack{} // Auxiliary stack `b`

	// Push the smallest element to `b`
	minIdx := findMinIndex(s)
	moveToTop(s, minIdx)
	fmt.Println("pb")
	b.Push(s.Pop())

	// If 5 elements, push another smallest one to `b`
	if len(s.elements) == 4 {
		minIdx = findMinIndex(s)
		moveToTop(s, minIdx)
		fmt.Println("pb")
		b.Push(s.Pop())
	}

	// Sort remaining 3 elements
	sortThree(s)

	// Push back elements from `b` to `a`
	for len(b.elements) > 0 {
		fmt.Println("pa")
		s.Push(b.Pop())
	}
}

// Moves the minimum value to the top
func moveToTop(s *Stack, index int) {
	if index == 0 {
		return
	} else if index <= len(s.elements)/2 {
		for i := 0; i < index; i++ {
			s.Rotate()
		}
	} else {
		for i := 0; i < len(s.elements)-index; i++ {
			s.ReverseRotate()
		}
	}
}

// Finds the index of the smallest element
func findMinIndex(s *Stack) int {
	minIdx, minVal := 0, s.elements[0]
	for i, val := range s.elements {
		if val < minVal {
			minVal = val
			minIdx = i
		}
	}
	return minIdx
}

// Main function to determine sorting strategy
func pushSwapSort(s *Stack) {
	switch len(s.elements) {
	case 2:
		sortTwo(s)
	case 3:
		sortThree(s)
	case 4, 5:
		sortFourOrFive(s)
	}
}

func main() {
	stack := &Stack{elements: []int{4, 1, 3, 2, 5}}
	fmt.Println("Initial Stack:", stack.elements)

	pushSwapSort(stack)

	fmt.Println("Sorted Stack:", stack.elements)
}
