package main

import (
	"fmt"
	"os"
	"strconv"
)

// declare the stack
type StackA struct {
	items []int
}

type StackB struct {
	items []int
}

func collectItems() []int {
	var result []int
	arg := os.Args[1:]
	for _, ar := range arg {
		num, _ := strconv.Atoi(ar)
		result = append(result, num)
	}
	return result
}

func (a *StackA) throwFirstTwo(b *StackB) (StackA, StackB) {
	var result StackA
	var res StackB
	for i := 0; i < len(a.items); i++ {
		if i > 1 {
			result.items = append(result.items, a.items[i])
		} else if i <= 1 {
			res.items = append([]int{a.items[i]}, res.items...)
		}
	}
	a = &result
	b = &res
	return *a, *b
}

// RadixSort sorts an array of integers using the radix sort algorithm.
func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum number to know the number of digits
	max := getMax(arr)

	// Perform counting sort for each digit
	for exp := 1; max/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}

	return arr
}

// getMax finds the maximum value in the array
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// countingSort performs a stable counting sort based on the digit at the given place (exp)
func countingSort(arr []int, exp int) []int {
	n := len(arr)
	output := make([]int, n) // Output array to store sorted numbers
	count := make([]int, 10) // Count array for digits (0-9)

	// Count occurrences of each digit in the current place (exp)
	for _, num := range arr {
		digit := (num / exp) % 10
		count[digit]++
	}

	// Update count[i] to store the position of the digit in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	return output
}

// CountingSort sorts an array of non-negative integers using the counting sort algorithm.
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum value in the array to determine the size of the count array
	max := getMax(arr)

	// Create a count array to store the frequency of each element
	count := make([]int, max+1)

	// Count the occurrences of each number in the array
	for _, num := range arr {
		count[num]++
	}

	// Build the sorted array
	index := 0
	for i, c := range count {
		for c > 0 {
			arr[index] = i
			index++
			c--
		}
	}

	return arr
}

// we use a maximum of 2 steps...using the radix sort algo...
func (a *StackA) solveThreeA() StackA {
	var a1 StackA

	a1.items = CountingSort(a.items)
	a = &a1
	return *a
}

func (a *StackA) push(b *StackB) (StackA, StackB) {
	var (
		a2 StackA
		a1 = *&a
		b1 = *&b
	)
	for i := range a.items {
		if i == 0 {
			b1.items = append([]int{a1.items[0]}, b1.items...)
		} else if i > 0 {
			a2.items = append(a2.items, a1.items[i])
		}
	}
	a = &a2
	b = *&b1
	return *a, *b
}

// implement a stack in an slice
func main() {
	a := StackA{collectItems()}
	var b StackB

	// throw the first 2 elements without checking to stack B...
	// r, i := a.throwFirstTwo(&b)
	// fmt.Println(r.solveThreeA(), i)
	r, i := a.push(&b)
	fmt.Println(r.push(&i))
}
