# _*Push-Swap*_

- The Push-Swap project involves sorting integers using a set of predefined stack operations. You are required to implement two programs in Go:
    - Push-swap: Calculates and displays the shortest sequence of instructions needed to sort a stack of integers.
    - Checker: Executes a sequence of instructions on a stack and verifies if the stack is correctly sorted.

## _*Instructions*_

- Program: push-swap
- The push-swap program takes a list of integers as arguments and outputs the smallest possible sequence of instructions to sort the stack.
Usage.

```sh
./push-swap "2 1 3 6 5 8"
```

### _*Expected Output*_

- The output will be a sequence of instructions to sort the stack, one per line. For example:

```bash
  pb
  pb
  ra
  sa
  rrr
  pa
  pa
```

### _*Error Handling*_

- If the arguments are invalid (e.g., non-integer values or duplicates), the program should output:

```bash
ERROR
```
