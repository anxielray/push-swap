package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return // No arguments given
	}
	instructions := pushSwap(os.Args[1])
	if instructions != "" {
		fmt.Print(instructions)
	}
}

// pushSwap sorts the stack and returns the instructions to do so.
func pushSwap(arg string) string {
	stackA := makeStackSlice(arg)
	if stackA == nil {
		return "Error\n"
	}

	if len(stackA) == 0 {
		return ""
	}

	sortedArray := make([]int, len(stackA))
	copy(sortedArray, stackA)
	sort.Ints(sortedArray)

	var stackB []int
	var instructions []string

	// Using simple cases for 2 or 3 elements
	switch len(stackA) {
	case 2:
		if stackA[0] > stackA[1] {
			instructions = append(instructions, "sa")
		}
	case 3:
		instructions = solveThree(&stackA)
	default:
		instructions = sortLargerStacks(&stackA, &stackB, sortedArray)
	}

	return strings.Join(instructions, "\n") + "\n"
}

// makeStackSlice converts the argument string to a slice of integers.
func makeStackSlice(arg string) []int {
	strArray := strings.Fields(arg)
	uniqueNumbers := make(map[int]bool)
	stack := []int{}

	for _, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil || uniqueNumbers[num] {
			fmt.Fprint(os.Stderr, "Error\n")
			return nil
		}
		uniqueNumbers[num] = true
		stack = append(stack, num)
	}
	return stack
}

// solveThree sorts a stack of three elements and returns the instructions.
func solveThree(stack *[]int) []string {
	instructions := []string{}
	if (*stack)[0] > (*stack)[1] {
		swap(stack)
		instructions = append(instructions, "sa")
	}
	if (*stack)[0] > (*stack)[2] {
		reverseRotate(stack)
		instructions = append(instructions, "rra")
	}
	if (*stack)[1] > (*stack)[2] {
		swap(stack)
		instructions = append(instructions, "sa")
	}
	return instructions
}

// sortLargerStacks implements sorting logic for larger stacks.
func sortLargerStacks(stackA *[]int, stackB *[]int, sortedArray []int) []string {
	instructions := []string{}
	for len(*stackA) > 0 {
		// Push elements to stack B in sorted order
		if len(*stackB) == 0 || (*stackA)[0] < (*stackB)[len(*stackB)-1] {
			instructions = append(instructions, "pb")
			*stackB = append([]int{(*stackA)[0]}, *stackB...)
			*stackA = (*stackA)[1:]
		} else {
			// Rotate or swap logic can be added here as needed
			instructions = append(instructions, "ra")
			rotate(stackA)
		}
	}

	// Move elements back from stack B to stack A
	for len(*stackB) > 0 {
		instructions = append(instructions, "pa")
		*stackA = append([]int{(*stackB)[0]}, *stackA...)
		*stackB = (*stackB)[1:]
	}

	return instructions
}

// swap swaps the first two elements of the stack.
func swap(stack *[]int) {
	if len(*stack) > 1 {
		(*stack)[0], (*stack)[1] = (*stack)[1], (*stack)[0]
	}
}

// rotate rotates the stack (first element becomes last).
func rotate(stack *[]int) {
	if len(*stack) > 0 {
		*stack = append((*stack)[1:], (*stack)[0])
	}
}

// reverseRotate reverses the rotation of the stack.
func reverseRotate(stack *[]int) {
	if len(*stack) > 0 {
		last := (*stack)[len(*stack)-1]
		*stack = append([]int{last}, (*stack)[:len(*stack)-1]...)
	}
}
