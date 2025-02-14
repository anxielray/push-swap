package main

import (
		"fmt"
		"math"
)

// Stack struct to simulate the push_swap stacks
type Stack struct {
	elements []int
	name     string
}

// Push an element to the stack
func (s *Stack) Push(val int) {
	s.elements = append([]int{val}, s.elements...)
	fmt.Printf("p%s\n", s.name)
}

// Pop an element from the stack
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
														return -1 // Stack empty
															}
																val := s.elements[0]
																	s.elements = s.elements[1:]
																		return val
																		}

																		// Swap top two elements
																		func (s *Stack) Swap() {
																				if len(s.elements) < 2 {
																							return
																								}
																									s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
																										fmt.Printf("s%s\n", s.name)
																										}

																										// Rotate upwards
																										func (s *Stack) Rotate() {
																												if len(s.elements) < 2 {
																															return
																																}
																																	first := s.elements[0]
																																		s.elements = append(s.elements[1:], first)
																																			fmt.Printf("r%s\n", s.name)
																																			}

																																			// Reverse rotate downwards
																																			func (s *Stack) ReverseRotate() {
																																					if len(s.elements) < 2 {
																																								return
																																									}
																																										last := s.elements[len(s.elements)-1]
																																											s.elements = append([]int{last}, s.elements[:len(s.elements)-1]...)
																																												fmt.Printf("rr%s\n", s.name)
																																												}

																																												// Find the index of the smallest element
																																												func findMinIndex(s *Stack) int {
																																														minIdx, minVal := 0, s.elements[0]
																																															for i, val := range s.elements {
																																																		if val < minVal {
																																																						minVal = val
																																																									minIdx = i
																																																											}
																																																												}
																																																													return minIdx
																																																													}

																																																													// Moves a given index to the top using the least number of moves
																																																													func moveToTop(s *Stack, index int) {
																																																															if index == 0 {
																																																																		return
																																																																			} else if index <= len(s.elements)/2 {
																																																																						for i := 0; i < index; i++ {
																																																																										s.Rotate()
																																																																												}
																																																																													} else {
																																																																																for i := 0; i < len(s.elements)-index; i++ {
																																																																																				s.ReverseRotate()
																																																																																						}
																																																																																							}
																																																																																							}

																																																																																							// Finds the best index to push from `a` to `b`
																																																																																							func findBestPushIndex(a *Stack) int {
																																																																																									bestIdx := 0
																																																																																										bestCost := math.MaxInt32

																																																																																											for i, val := range a.elements {
																																																																																														cost := i // Rotating up
																																																																																																if len(a.elements)-i < cost { // Reverse rotating down
																																																																																																			cost = len(a.elements) - i
																																																																																																					}

																																																																																																							if cost < bestCost {
																																																																																																											bestCost = cost
																																																																																																														bestIdx = i
																																																																																																																}
																																																																																																																	}
																																																																																																																		return bestIdx
																																																																																																																		}

																																																																																																																		// Greedy sorting for medium-sized stacks
																																																																																																																		func greedyPushSwap(a *Stack) {
																																																																																																																				b := &Stack{name: "b"}

																																																																																																																					// Push elements to `b` (except 3)
																																																																																																																						for len(a.elements) > 3 {
																																																																																																																									bestIdx := findBestPushIndex(a)
																																																																																																																											moveToTop(a, bestIdx)
																																																																																																																													b.Push(a.Pop())
																																																																																																																														}

																																																																																																																															// Sort remaining 3 elements
																																																																																																																																sortThree(a)

																																																																																																																																	// Push back elements from `b` to `a`
																																																																																																																																		for len(b.elements) > 0 {
																																																																																																																																					moveToInsertCorrectly(a, b.Pop())
																																																																																																																																						}
																																																																																																																																						}

																																																																																																																																						// Sorts a stack of exactly 3 elements optimally
																																																																																																																																						func sortThree(s *Stack) {
																																																																																																																																								a, b, c := s.elements[0], s.elements[1], s.elements[2]
																																																																																																																																									if a > b && b > c { // 321
																																																																																																																																											s.Swap()
																																																																																																																																													s.ReverseRotate()
																																																																																																																																														} else if a > c && c > b { // 312
																																																																																																																																																s.Rotate()
																																																																																																																																																	} else if b > a && a > c { // 231
																																																																																																																																																			s.ReverseRotate()
																																																																																																																																																				} else if b > c && c > a { // 213
																																																																																																																																																						s.Swap()
																																																																																																																																																							} else if c > a && a > b { // 132
																																																																																																																																																									s.Swap()
																																																																																																																																																											s.Rotate()
																																																																																																																																																												}
																																																																																																																																																												}

																																																																																																																																																												// Inserts value from `b` into the correct place in `a`
																																																																																																																																																												func moveToInsertCorrectly(a *Stack, value int) {
																																																																																																																																																														pos := 0
																																																																																																																																																															for i, val := range a.elements {
																																																																																																																																																																		if value < val {
																																																																																																																																																																						pos = i
																																																																																																																																																																									break
																																																																																																																																																																											}
																																																																																																																																																																												}

																																																																																																																																																																													moveToTop(a, pos)
																																																																																																																																																																														a.Push(value)
																																																																																																																																																																														}

																																																																																																																																																																														// Main function to test the implementation
																																																																																																																																																																														func main() {
																																																																																																																																																																																stack := &Stack{elements: []int{8, 3, 7, 2, 6, 1, 4, 5}, name: "a"}
																																																																																																																																																																																	fmt.Println("Initial Stack:", stack.elements)

																																																																																																																																																																																		greedyPushSwap(stack)

																																																																																																																																																																																			fmt.Println("Sorted Stack:", stack.elements)
																																																																																																																																																																																			}
																																																																																																																																																																														}
																																																																																																																																																																		}
																																																																																																																																																															}
																																																																																																																																																												}
																																																																																																																																						}
																																																																																																																																		}
																																																																																																																						}
																																																																																																																		}
																																																																																																							}
																																																																																											}
																																																																																							}
																																																																																}
																																																																													}
																																																																						}
																																																																			}
																																																															}
																																																													}
																																																		}
																																															}
																																												}
																																					}
																																			}
																												}
																										}
																				}
																		}
											}
									}
						}
			}
)
