package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    arguments()
    push_swap(os.Args[1]) // prints out the required commands to sort the stack elements.
    // now the algorithm used to run the program...
}

// grams function
func push_swap(arg string) {
    sortedArray := makeStackSlice(arg)
    sort.Ints(sortedArray)
    var (
        stackA       []int
        stackB       []int
        instructions string
    )
    stackA = makeStackSlice(arg)

    // check if the first 2 elements can be pushed to stack b
    if validateFirstPush(stackA, sortedArray) {
        stackA, stackB, instructions = firstPush(stackA, stackB)
        instructions += "pb
pb
"
    } else if validateRotate(stackA, sortedArray) {
        stackA = rotate(stackA)
        stackA, stackB = push(stackA, stackB)
        instructions += "ra
pb
"
    } else if validateReverseRotate(stackA, sortedArray) {
        stackA = reverseRotate(stackA)
        stackA, stackB = push(stackA, stackB)
        instructions += "rra
pb
"
    }

    // consider arranging the rest of the immigrants to b in descendin order
    if validateRRR(stackA, stackB, sortedArray) {
        stackA, stackB = bothReverseRotate(stackA, stackB)
        stackA, stackB = push(stackA, stackB)
        instructions += "rrr
pb
"
    }

    for i := range sortedArray {
        if len(stackA) == 3 && (stackA[0] != sortedArray[i] && stackA[1] != stackA[(i+1)] && stackA[1] != sortedArray[i+2]) {
            stackA, instructions = solveThree(stackA)
            instructions += ""
        }
    }
    var (
        test []int
        flag bool
    )
    test = append(test, stackB[(len(stackB)-1):]...)
    test = append(test, stackA...)
    for i, num := range test {
        for x, n := range sortedArray {
            if (i == x) && (num == n) {
                flag = true
            }
        }
    }
    if flag {
        fmt.Println(instructions)
    }
}

// the function will take an argument string and apppednd each number to a slice of int.
func makeStackSlice(arg string) (sorted []int) {
    array := strings.Fields(arg)
    for _, str := range array {
        num, err := strconv.Atoi(str)
        if err != nil {
            fmt.Println("Error")
            os.Exit(0)
        }
        sorted = append(sorted, num)
    }
    return
}

// INSTRUCTIONS:
// swapping the first two elements
func swap(array []int) (swapped []int) {
    swapped = append(array[2:], swapped...)
    swapped = append([]int{array[0]}, swapped...)
    swapped = append([]int{array[1]}, swapped...)
    return
}

// rotate(the first element becomes last)
func rotate(array []int) (rotated []int) {
    rotated = append(array[1:], rotated...)
    rotated = append(rotated, array[0])
    return
}

// reverse rotate(the last elemen becomes first)
func reverseRotate(array []int) (reverseRotated []int) {
    reverseRotated = append(array[:len(array)-1], reverseRotated...)
    reverseRotated = append([]int{array[len(array)-1]}, reverseRotated...)
    return
}

// RRR command
func bothReverseRotate(stackA, stackB []int) (rrrA, rrrB []int) {
    return reverseRotate(stackA), reverseRotate(stackB)
}

// SS command
func bothSwap(stackA, stackB []int) (ssA, ssB []int) {
    return swap(stackA), swap(stackB)
}

// RR command
func bothRotate(stackA, stackB []int) (rrA, rrB []int) {
    return rotate(stackA), rotate(stackB)
}

// push command (should add an element to a host array and delete it from previous host array)
func push(previousHost, newHost []int) (previous, current []int) {
    current = append([]int{previousHost[0]}, newHost...)
    previous = previousHost[1:]
    return
}

// ALGORITHM:
// check if the first two elements do not fall in the last half of the sorted array
func validateFirstPush(stackA, sortedArray []int) bool {
    if stackA[0] < sortedArray[len(stackA)/2] {
        if stackA[1] < sortedArray[len(stackA)/2] {
            return true
        }
    }
    return false
}

// incase the validate first throws a false, we validate for a reverse rotate or a rotate.
func validateRotate(stackA, sortedArray []int) bool {
    if stackA[0] >= sortedArray[len(stackA)/2] && stackA[1] < sortedArray[len(stackA)/2] {
        return true
    }
    return false
}

// validate a reverse rotate
func validateReverseRotate(stackA, sortedArray []int) bool {
    if stackA[0] >= sortedArray[len(stackA)/2] && stackA[len(stackA)-1] < sortedArray[len(stackA)/2] {
        return true
    }
    return false
}

// validate reverse rotate
func validateRRR(stackA, stackB, sortedArray []int) bool {
    if validateReverseRotate(stackA, sortedArray) && validateReverseRotate(stackB, sortedArray) {
        return true
    }
    return false
}

// function to solve the 3 elements command
func solveThree(stack []int) (threeSolved []int, instructions string) {
    if stack[0] > stack[1] && stack[0] > stack[2] {
        rotated := rotate(stack)
        instructions += "ra
"
        if rotated[0] > rotated[1] {
            threeSolved = swap(rotated)
            instructions += "sa
"
        } else {
            threeSolved = rotated
        }
    } else if stack[0] > stack[1] && stack[0] < stack[2] {
        threeSolved = swap(stack)
        instructions += "sa
"
    } else if stack[0] < stack[1] && stack[0] > stack[2] {
        threeSolved = reverseRotate(stack)
        instructions += "rra
"
    }
    return
}

// throw the first 2 elements to stack b regardless
func firstPush(stackA, stackB []int) (a, b []int, instruction string) {
    for x := 1; x <= 2; x++ {
        a, b = push(stackA, stackB)
        instruction += "pb
"
    }
    return
}

// check for the number of arguments passed on the CMD
func arguments() {
    if len(os.Args) != 2 {
        return
    }
}

