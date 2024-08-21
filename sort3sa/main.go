package main

import "fmt"

// sortThree sorts a slice of three integers in ascending or descending order
// using a maximum of two swaps.
func sortThreeA(stack []int) (instruction string) {
	// for stack a, the slice should be in descending order
	// check if it is sorted for stack a(base case)
	if stack[0] < stack[1] && stack[1] < stack[2] {
		instruction = ""
	}
	// check for the situations of [1, 3, 2]
	if stack[0] < stack[1] && stack[1] > stack[2] {
		instruction = "rra\nsa\n"
	}
	// check for the [2, 1, 3]
	if stack[0] > stack[1] && stack[0] < stack[2] {
		instruction = "sa\n"
	}
	// check for the [2, 3, 1]
	if stack[0] > stack[2] && stack[0] < stack[1] {
		instruction = "rra\n"
	}
	// check for [3, 1, 2]
	if stack[0] > stack[1] && stack[0] > stack[2] && stack[1] < stack[2] {
		instruction = "rra\nrra\n"
	}
	// check for [3, 2, 1]
	if stack[0] > stack[1] && stack[1] > stack[2] {
		instruction = "sa\nrra\n"
	}
	return
}

func main() {
	fmt.Println(sortThreeA([]int{3, 2, 1}))
	println("sorted")
}
