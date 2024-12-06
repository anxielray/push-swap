package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// declare the variable that  will hold the commands
var commands []string

// declare the stack
type StackA struct {
	items []int
}

type StackB struct {
	items []int
}

func CollectItems() []int {
	var result []int
	arg := os.Args[1:]
	for _, ar := range arg {
		num, _ := strconv.Atoi(ar)
		result = append(result, num)
	}
	return result
}

// RadixSort sorts an array of integers using the radix sort algorithm.
func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum number to know the number of digits
	max := GetMax(arr)

	// Perform counting sort for each digit
	for exp := 1; max/exp > 0; exp *= 10 {
		arr = countingSort(arr, exp)
	}

	return arr
}

// getMax finds the maximum value in the array
func GetMax(arr []int) int {
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
	max := GetMax(arr)

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

func main() {

	var (
		a StackA
		b StackB
	)
	a.items = CollectItems()

	//create a  sorted array
	var sorted = RadixSort(CollectItems())

	//identify the median element in  the array
	med := FindMedian(sorted)

	var count int
	//push to b what is less than the median element
	a, b, count = a.MedPush(med, &b)

	a = a.SwapA()
	b = b.SwapB()

	//check for the validations of the commands ss, rr and rrr
	commands = Rr(Rr(Ss(commands)))

	//push back elements to a
	a, b = b.PushBack(count, &a)

	//print the commannds
	for _, co := range commands {
		fmt.Println(co)
	}

}

func (b *StackB) PushBack(n int, a *StackA) (StackA, StackB) {
	for i := 0; i < n; i++ {
		*a, *b = b.PushA(a)
	}
	return *a, *b
}

func (a *StackA) SwapA() StackA {

	//the key will be the element and the value will be the index in the array
	originalMp := make(map[int]int)
	for i, n := range a.items {
		originalMp[n] = i
	}

	//the key will be the element and the value will be the new index in the sorted array
	sortedMp := make(map[int]int)
	sortd := RadixSort(a.items)
	for i, n := range sortd {
		sortedMp[n] = i
	}

	//compare the values of one map and the keys of another and vice versa
	for _, v := range originalMp {
		for _, v1 := range sortedMp {
			if sortedMp[a.items[0]] == originalMp[a.items[1]] && v == v1 {
				a.items[0], a.items[1] = a.items[1], a.items[0]
				commands = append(commands, "sa")
				break
			}
		}
		break
	}
	return *a
}

func (b *StackB) SwapB() StackB {

	//the key will be the element and the value will be the index in the array
	originalMp := make(map[int]int)
	for i, n := range b.items {
		originalMp[n] = i
	}

	//the key will be the element and the value will be the new index in the sorted array
	sortedMp := make(map[int]int)
	var sortd []int
	for i := len(RadixSort(b.items)) - 1; i >= 0; i-- {
		sortd = append(sortd, RadixSort(b.items)[i])
	}

	for i, n := range sortd {
		sortedMp[n] = i
	}

	//compare the values of one map and the keys of another and vice versa
	for _, v := range originalMp {
		for _, v1 := range sortedMp {
			if sortedMp[b.items[0]] == originalMp[b.items[1]] && v == v1 {
				b.items[0], b.items[1] = b.items[1], b.items[0]
				commands = append(commands, "sb")
				break
			}
		}
		break
	}
	return *b
}

func (a *StackA) PushB(b *StackB) (StackA, StackB) {

	var (
		a1 StackA
		b1 StackB
	)

	b1.items = append(b1.items, b.items...)
	for i, n := range a.items {
		if i != 0 {
			a1.items = append(a1.items, n)
		} else {
			b1.items = append([]int{n}, b1.items...)
		}
	}
	a = &a1
	b = &b1
	commands = append(commands, "pb")
	return *a, *b
}

func (b *StackB) PushA(a *StackA) (StackA, StackB) {

	var (
		a1 StackA
		b1 StackB
	)

	a1.items = append(a1.items, a.items...)
	for i, n := range b.items {
		if i != 0 {
			b1.items = append(b1.items, n)
		} else {
			a1.items = append([]int{n}, a1.items...)
		}
	}
	a = &a1
	b = &b1
	commands = append(commands, "pa")
	return *a, *b
}

func (a *StackA) RotateA() StackA {

	var a1 StackA
	for i := range a.items {
		if i > 0 {
			a1.items = append(a1.items, a.items[i])
		}
	}
	a1.items = append(a1.items, a.items[0])
	a = &a1
	commands = append(commands, "ra")
	return *a
}

func (b *StackB) RotateB() StackB {

	var b1 StackB
	for i := range b.items {
		if i > 0 {
			b1.items = append(b1.items, b.items[i])
		}
	}
	b1.items = append(b1.items, b.items[0])
	b = &b1
	commands = append(commands, "rb")
	return *b
}

func (a *StackA) MedPush(med int, b *StackB) (StackA, StackB, int) {

	var (
		a1    StackA
		b1    StackB
		count int
	)
	a1 = *a
	b1 = *b

	//if the value is less than the median number, we are allowed to push it to b
	for _, n := range a.items {
		if n < med {
			//push to b
			a1, b1 = a1.PushB(&b1)
			count++
		} else {
			//rotate
			a1 = a1.RotateA()
		}
	}
	a, b = &a1, &b1
	return *a, *b, count
}

func FindMedian(a []int) int {
	l := float64(len(a))
	return a[int(math.Floor(l/2.0))]
}

func Ss(commands []string) []string {

	var result []string
	for i := 0; i < len(commands)-1; i++ {
		c := commands[i]
		if c == "sa" {
			if commands[i+1] == "sb" {
				result = append(result, "ss")
				i++
				continue
			} else {
				result = append(result, c)
			}
		} else if c == "sb" {
			if commands[i+1] == "sa" {
				result = append(result, "ss")
				i++
				continue
			}
		} else {
			result = append(result, c)
		}

	}
	if commands[len(commands)-2] == "sb" {
		if commands[len(commands)-1] != "sa" {
			result = append(result, commands[len(commands)-1])
		}
	} else if commands[len(commands)-2] == "sa" {
		if commands[len(commands)-1] != "sb" {
			result = append(result, commands[len(commands)-1])
		}
	} else {
		result = append(result, commands[len(commands)-1])
	}
	commands = result
	return commands
}

func Rr(commands []string) []string {

	var result []string
	for i := 0; i < len(commands)-1; i++ {
		c := commands[i]
		if c == "ra" {
			if commands[i+1] == "rb" {
				result = append(result, "rr")
				i++
				continue
			} else {
				result = append(result, c)
			}
		} else if c == "rb" {
			if commands[i+1] == "ra" {
				result = append(result, "rr")
				i++
				continue
			}
		} else {
			result = append(result, c)
		}

	}
	result = append(result, commands[len(commands)-1])
	commands = result
	return commands
}

func Rrr(commands []string) []string {

	var result []string
	for i := 0; i < len(commands)-1; i++ {
		c := commands[i]
		if c == "rra" {
			if commands[i+1] == "rrb" {
				result = append(result, "rrr")
				i++
				continue
			}
		} else if c == "rrb" {
			if commands[i+1] == "rra" {
				result = append(result, "rrr")
				i++
				continue
			} else {
				result = append(result, c)
			}
		} else {
			result = append(result, c)
		}

	}
	result = append(result, commands[len(commands)-1])
	commands = result
	return commands
}
