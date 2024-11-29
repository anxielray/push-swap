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
func (a *StackA) solveThreeA() *StackA {
	var a1 StackA

	a1.items = CountingSort(a.items)
	a = &a1
	return a
}

func (a *StackA) pushB(b *StackB) (StackA, StackB) {
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
	fmt.Println("pb")
	return *a, *b
}

func (b *StackB) pushA(a *StackA) (StackA, StackB) {
	var (
		a1 StackA
		b1 StackB
	)
	a1.items = append(a1.items, a.items...)
	for i := range b.items {
		if i != 0 {
			b1.items = append(b1.items, b.items[i])
		} else {
			a1.items = append([]int{b.items[i]}, a1.items...)
		}
	}

	b = &b1
	a = &a1

	fmt.Println("pa")
	return *a, *b
}

// implement a stack in an slice
func main() {
	a := StackA{collectItems()}
	var b StackB

	// throw the first 2 elements without checking to stack B...
	// r, i := a.throwFirstTwo(&b)
	r, i := a.pushB(&b)
	fmt.Println(r, i)
	x, q := r.pushB(&i)
	fmt.Println(x, q)
	x = *x.solveThreeA()

	//sort the elements in stack b in descending order...(We will later work on tracking steps taken to achieve this)
	q.items = reverseSlice(RadixSort(q.items))
	fmt.Println("ss")
	fmt.Println(x, q)

	//check for the target element
	targetElementA1 := x.TargetElement(q.items[0])

	//mark the element at index 0
	// startIndex := x.items[0]

	//bring the target element to the top
	z := *x.TargetElementToTheTop(targetElementA1)
	fmt.Println(z, q)

	//push back an element to stack a
	c, d := q.pushA(&z)

	//re-organize the elements in stack a
	g := *(c.RotateA()).RotateA()

	targetElementA2 := g.TargetElement(d.items[0])

	//bring the target element to the top
	h := *(g.TargetElementToTheTop(targetElementA2))
	h = *h.TargetElementToTheTop(g.TargetElement(d.items[0]))
	fmt.Println(h, d)

	//push back another element to stack b
	e, f := d.pushA(&h)
	fmt.Println(e, f)

	n := *((e.RotateA()).RotateA()).RotateA()
	fmt.Println(n, f)
}

func (a *StackA) RotateA() *StackA {
	var (
		a1 StackA
	)

	for i := range a.items {
		if i > 0 {
			a1.items = append(a1.items, a.items[i])
		}
	}
	a1.items = append(a1.items, a.items[0])
	a = &a1
	fmt.Println("ra")
	return a
}

func HowManyFromTheTop(a []int, top int) int {
	var aim int
	for i, r := range a {
		if r == top {
			aim = i
			break
		}
	}
	return aim + 1
}

func reverseSlice(a []int) []int {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func (a *StackA) TargetElement(elem int) int {
	var test StackA
	test.items = append(test.items, a.items...)
	test.items = append(test.items, elem)
	temp := RadixSort(test.items)
	return getIndex(temp, elem) + 1
}

func getIndex(a []int, target int) int {
	for i, r := range a {
		if r == target {
			return i
		}
	}
	return -1
}

func (a *StackA) TargetElementToTheTop(targetIndex int) *StackA {

	//check the length of the stack A
	l := len(a.items)
	diff := l - targetIndex
	if diff == 0 {
		fmt.Println("rra")
	}
	a = a.ReverseRotate()
	return a
}

func (a *StackA) ReverseRotate() *StackA {
	var (
		a1 StackA
	)

	for i, r := range a.items {
		if i < len(a.items)-1 {
			a1.items = append(a1.items, r)
		}
	}
	a1.items = append([]int{a.items[len(a.items)-1]}, a1.items...)
	a = &a1
	return a
}
