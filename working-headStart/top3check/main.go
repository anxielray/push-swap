package main

import (
	"fmt"
	"os"
	"sort"
)

// isAnyTopThree checks if any of the provided integers are among the top 3 largest elements
// in the sorted slice.
func isAnyTopThree(slice []int, a, b, c int) bool {
	if len(slice) < 3 {
		println("The slice must contain at least three elements.")
		os.Exit(0)
	}

	// Sort the slice in descending order
	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedSlice)))

	// Get the top 3 largest elements
	topThree := sortedSlice[:3]

	// Check if any of the given integers are in the top three
	for _, value := range []int{a, b, c} {
		for _, top := range topThree {
			if value == top {
				return true
			}
		}
	}

	return false
}

func main() {
	// Example usage
	slice := []int{4, 8, 1, 6, 3, 9, 10, 0}
	a, b, c := 6, 2, 5

	result := isAnyTopThree(slice, a, b, c)
	fmt.Println(result) // Output: true (since 8 is among the top 3 largest elements)

	d, e, f := 7, 10, 12
	result = isAnyTopThree(slice, d, e, f)
	fmt.Println(result) // Output: false (since none of 7, 10, or 12 are among the top 3 largest elements)
}
