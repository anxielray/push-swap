package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//declare the stack a
type Stack struct {
	elements []int
}

//the push command
func (s *Stack) push(value int) {
	s.elements = append([]int{value}, s.elements...)
}

func (s *Stack) pop() int {
	if len(s.elements) == 0 {
		return 0
	}
	value := s.elements[0]
	s.elements = s.elements[1:]
	return value
}

func (s *Stack) top() int {
	if len(s.elements) == 0 {
		return 0
	}
	return s.elements[0]
}

func (s *Stack) swap() {
	if len(s.elements) < 2 {
		return
	}
	s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
}

func (s *Stack) rotate() {
	if len(s.elements) < 2 {
		return
	}
	first := s.elements[0]
	s.elements = append(s.elements[1:], first)
}

func (s *Stack) reverseRotate() {
	if len(s.elements) < 2 {
		return
	}
	last := s.elements[len(s.elements)-1]
	s.elements = append([]int{last}, s.elements[:len(s.elements)-1]...)
}

func executeInstructions(a, b *Stack, instructions []string) {
	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			a.swap()
		case "sb":
			b.swap()
		case "ss":
			a.swap()
			b.swap()
		case "pa":
			a.push(b.pop())
		case "pb":
			b.push(a.pop())
		case "ra":
			a.rotate()
		case "rb":
			b.rotate()
		case "rr":
			a.rotate()
			b.rotate()
		case "rra":
			a.reverseRotate()
		case "rrb":
			b.reverseRotate()
		case "rrr":
			a.reverseRotate()
			b.reverseRotate()
		default:
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
}

func isSorted(stack *Stack) bool {
	for i := 1; i < len(stack.elements); i++ {
		if stack.elements[i] < stack.elements[i-1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	a := &Stack{}
	b := &Stack{}

	args := os.Args[1:]
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.push(num)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		instructions = append(instructions, line)
	}

	executeInstructions(a, b, instructions)

	if isSorted(a) && len(b.elements) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

