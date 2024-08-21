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
		stck.reverse_rotate()
		stck.swap()
		instruction = "rra\nsa\n"
	}
	// check for the [2, 1, 3]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] < (*stck)[2] {
		stck.swap()
		instruction = "sa\n"
	}
	// check for the [2, 3, 1]
	if (*stck)[0] > (*stck)[2] && (*stck)[0] < (*stck)[1] {
		stck.reverse_rotate()
		instruction = "rra\n"
	}
	// check for [3, 1, 2]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] > (*stck)[2] && (*stck)[1] < (*stck)[2] {
		stck.reverse_rotate()
		stck.reverse_rotate()
		instruction = "rra\nrra\n"
	}
	// check for [3, 2, 1]
	if (*stck)[0] > (*stck)[1] && (*stck)[1] > (*stck)[2] {
		stck.swap()
		stck.reverse_rotate()
		instruction = "sa\nrra\n"
	}
	return
}

// sortThree sorts a slice of three integers in ascending or descending order
// using a maximum of two swaps.
func sortThreeB(stck *stack) (instruction string) {
	// for (*stck) b, the slice should be in descending order
	// check if it is sorted for (*stck) b
	if (*stck)[0] > (*stck)[1] && (*stck)[1] > (*stck)[2] {
		instruction = ""
	}
	// check for the situations of [1, 3, 2] (for (*stck) b, we will need to rrb, rrb)
	if (*stck)[0] < (*stck)[1] && (*stck)[1] > (*stck)[2] {
		instruction = "rrb\nrrb\n"
	}
	// check for the [2, 1, 3]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] < (*stck)[2] {
		instruction = "sb\n"
	}
	// check for the [2, 3, 1]
	if (*stck)[0] > (*stck)[2] && (*stck)[0] < (*stck)[1] {
		instruction = "rrb\n"
	}
	// check for [3, 1, 2]
	if (*stck)[0] > (*stck)[1] && (*stck)[0] > (*stck)[2] && (*stck)[1] < (*stck)[2] {
		instruction = "rrb\nsb\n"
	}
	// check for  [1, 2, 3]
	if (*stck)[0] < (*stck)[1] && (*stck)[1] < (*stck)[2] {
		instruction = "sb\nrrb"
	}
	return
}

// sorts
func push_swap(a, b *stack) (instruction string) {
	// make sure we have something to compare in the stack
	if len(*a) <= 2 {
		println("The number of arguments you provided are not enough, add one more...")
		os.Exit(0)
	}

	// working with the staack of length more than 2 (3 > ...)
	for len(*a) > 2 {
		// check if the stack came with only 3 elements inside...
		if len(*a) != 3 {

			// convert the string passed in the command line as a slice of string to be used by other functions thst need it in that form...
			slice := parseIntSlice(os.Args[1])

			// Since we have to start by pushing to stack b, confirm if we have the least elements in the sorted stack at the very bottom of the stack
			if isAnyTopThreeSmallest(slice, slice[len(slice)-1]) {
				a.reverse_rotate()
				instruction += "rra\n"
			} else if isAnyTopThreeSmallest(slice, slice[len(slice)-1]) {
				a.reverse_rotate() // first roation
				a.reverse_rotate() // second rotation
				instruction += "rra\nrra\n"
			} else if isAnyTopThree(slice, slice[0]) && !isAnyTopThree(slice, slice[1]) { // find out if the top 2 elements, any of them is the amongst the 3 largest elements in the sorted stack
				a.rotate()
				instruction += "ra\n"
			} else if isAnyTopThree(slice, slice[0]) && isAnyTopThree(slice, slice[1]) {
				a.rotate()
				a.rotate()
				instruction += "ra\nra\n"
			}

			// do the first 2 pushes to stack b regardless...
			b.push(a.pop())
			b.push(a.pop())
			instruction += "pb\npb\n"

			// check if the arramgement is correct for the incoming 3rd number...
			if (*b)[0] < (*b)[1] {
				b.swap()
				instruction += "sb\n"
			}

			// perform the 3rd push to stack b
			b.push(a.pop())
			instruction += "pb\n"
			continue
		} else if len(*a) == 3 && len(*b) == 3 { // if the elements inside of stack a is only 3, then, the function to stack 3 elements is called
			instruction += sortThreeA(a)
			instruction += sortThreeB(b)
		}
	}

	// push back to a till stack b becomes empty...
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
	// confrim if the arguments are 2 on the command line; the name of the program and the argument...
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

func isAnyTopThreeSmallest(slice []int, a int) bool {
	if len(slice) < 3 {
		panic("The slice must contain at least three elements.")
	}

	// Sort the slice in ascending order
	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Ints(sortedSlice)

	// Get the top 3 smallest elements
	topThreeSmallest := sortedSlice[:3]

	// Check if any of the given integers are in the top three smallest
	for _, top := range topThreeSmallest {
		if a == top {
			return true
		}
	}

	return false
}

// isSorted compares the original slice to its sorted version.
// It returns true if the original slice is already sorted in ascending order, otherwise false.
func isSorted(slice *[]int) bool {
	// Create a copy of the original slice
	sortedSlice := make([]int, len(*slice))
	copy(sortedSlice, *slice)

	// Sort the copy
	sort.Ints(sortedSlice)

	// Compare the original slice to the sorted copy
	for i := range *slice {
		if (*slice)[i] != sortedSlice[i] {
			return false
		}
	}

	return true
}
