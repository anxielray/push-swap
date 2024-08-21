package main

import (
	"fmt"
	"os"
	"sort"
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

// sorts 3 elements left in stack a
func sortThreeA(stck *stack) (instruction string) {
	// for stack a, the slice should be in descending order
	// check if it is sorted for stack a(base case)
	if (*stck)[0] < (*stck)[1] && (*stck)[1] < (*stck)[2] {
		instruction = ""
	}
	// check for the situations of [1, 3, 2]
	if (*stck)[0] < (*stck)[1] && (*stck)[1] > (*stck)[2] {
		instruction = "rra\nsa\n"
	}
	// check for the [2, 1, 3]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] < (*stck)[2] {
		instruction = "sa\n"
	}
	// check for the [2, 3, 1]
	if (*stck)[0] > (*stck)[2] && (*stck)[0] < (*stck)[1] {
		instruction = "rra\n"
	}
	// check for [3, 1, 2]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] > (*stck)[2] && (*stck)[1] < (*stck)[2] {
		instruction = "rra\nrra\n"
	}
	// check for [3, 2, 1]
	if (*stck)[0] > (*stck)[1] && (*stck)[1] > (*stck)[2] {
		instruction = "sa\nrra\n"
	}
	return
}

// sorts
func push_swap(a, b *stack) (instruction string) {
	// make sure we have something to compare in the stack
	if len(*a) < 2 {
		return
	}

	for len(*a) > 2 {
		// start by confirming if we have the correct count of elements in the stack. If the stack has only 3 elements...
		if len(*a) != 3 {
			//Since we have to start by pushing to stack b, confirm if we have the least elements in the sorted stack at the very bottom of the stack
			
			// find out if the top 2 elements, any of them is the amongst the 3 largest elements in the sorted stack
			slice := parseIntSlice(os.Args[1])
			if isAnyTopThree(slice, slice[0])&& !isAnyTopThree(slice, slice[1]) {
				instruction += "ra\n"
			}else if isAnyTopThree(slice, slice[1]) && isAnyTopThree(slice, slice[1]) {
				a.rotate()
				a.rotate()
				instruction += "ra\nra\n"
			}else {
				instruction += 
			}
			

			if (*a)[0] < (*a)[1] {
				b.push(a.pop())
				instruction += "pb\n"
			} else { // I generally feel like this is part is irrelevant
				a.rotate()
				instruction += "ra\n"
			}
		} else if len(*a) == 3 {
			instruction += sortThreeA(a)
		}
	}

	for len(*b) > 0 {
		a.push(b.pop())
		instruction += "pa\n"
	}

	if (*a)[0] > (*a)[1] {
		a.swap()
		instruction += "sa\n"
	}
	return
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
	fmt.Println(push_swap(&a, &b))
}

// isAnyTopThree checks if any of the provided integers are among the top 3 largest elements
// in the sorted slice.
func isAnyTopThree(stack []int, a int) bool {
	if len(stack) < 3 {
		println("The slice must contain at least three elements.")
		os.Exit(0)
	}

	// Sort the slice in descending order
	sortedSlice := make([]int, len(stack))
	copy(sortedSlice, stack)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSlice)))

	// Get the top 3 largest elements
	topThree := sortedSlice[:3]

	// Check if any of the given integers are in the top three
	for _, top := range topThree {
		if a == top {
			return true
		}
	}

	return false
}

func parseIntSlice(s string) (slice []int) {
	// Split the input string by spaces
	strSlice := strings.Split(s, " ")

	// Create a slice to hold the integers
	intSlice := make([]int, len(strSlice))

	// Convert each string to an integer and store it in the intSlice
	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			println("Error parsing string to int") // Return the error if the conversion fails
			os.Exit(0)
		}
		intSlice[i] = num
	}

	return intSlice
}
