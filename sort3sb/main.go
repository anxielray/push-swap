package main

import "fmt"

// sortThree sorts a slice of three integers in ascending or descending order
// using a maximum of two swaps.
func sortThreeB(stack []int) (instruction string) {
	// for stack b, the slice should be in descending order
	// check if it is sorted for stack b
	if stack[0] > stack[1] && stack[1] > stack[2] {
		instruction = ""
	}
	// check for the situations of [1, 3, 2] (for stack b, we will need to rrb, rrb)
	if stack[0] < stack[1] && stack[1] > stack[2] {
		instruction = "rrb\nrrb\n"
	}
	// check for the [2, 1, 3]
	if stack[0] > stack[1] && stack[0] < stack[2] {
		instruction = "sb\n"
	}
	// check for the [2, 3, 1]
	if stack[0] > stack[2] && stack[0] < stack[1] {
		instruction = "rrb\n"
	}
	// check for [3, 1, 2]
	if stack[0] > stack[1] && stack[0] > stack[2] && stack[1] < stack[2] {
		instruction = "rrb\nsb\n"
	}
	return
}

func main() {
	fmt.Println(sortThreeB([]int{3, 2, 1}))
	println("sorted")
}
