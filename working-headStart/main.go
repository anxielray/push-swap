package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Initialize the stack a as a common slice to be shared through out the program
type stack []int

// state a method to push an element to another stack
func (s *stack) push(x int) {
	*s = append(*s, x)
}

// deletes an element from a stack after pushing it to another stack...(I feel like there should be a review in this method)
func (s *stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

// Swap the first two elements of a stack
func (s *stack) swap() {
	(*s)[0], (*s)[1] = (*s)[1], (*s)[0]
}

// take the top element to the bottom of the stack
func (s *stack) rotate() {
	(*s)[0], *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
}

// take the last element to the top of the stack
func (s *stack) reverse_rotate() {
	*s = append([]int{(*s)[len(*s)-1]}, (*s)[:len(*s)-1]...)
}

func push_swap(a, b *stack) {

	//make sure we have something to compare in the stack
	if len(*a) < 2 {
		return
	}

	//if the top most element is greater than the second element, swap the stack
	if (*a)[0] > (*a)[1] {
		a.swap()
		fmt.Println("sa")
	}

	for len(*a) > 2 {

		//after confirming the order to which the first 2 elements, push them to stack b

		if (*a)[0] > (*a)[1] && (*a)[1] > (*a)[2] {
			// Case 1: 3 2 1 - reverse order
			a.swap() // Swap top two: 2 3 1
			fmt.Println("sa")
			a.reverse_rotate() // Rotate to get: 1 2 3
			fmt.Println("rra")
		}

		if (*a)[0] < (*a)[1] {
			b.push(a.pop())
			fmt.Println("pb")
		} else { //I generally feel like this is part is irrelevant
			a.rotate()
			fmt.Println("ra")
		}
	}

	for len(*b) > 0 {
		a.push(b.pop())
		fmt.Println("pa")
	}

	if (*a)[0] > (*a)[1] {
		a.swap()
		fmt.Println("sa")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run push_swap.go \"2 1 3 6 5 8\"")
		return
	}

	var a stack
	for _, arg := range strings.Fields(os.Args[1]) {
		x, _ := strconv.Atoi(arg)
		a.push(x)
	}

	b := make(stack, 0)
	push_swap(&a, &b)
}
