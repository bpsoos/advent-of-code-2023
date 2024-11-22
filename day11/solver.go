package day11

import (
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day11_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 11 test solution 1")
	fmt.Println(testInput)
	fmt.Println("day 11 test solution 2")
	fmt.Println()

	input, err := os.ReadFile("inputs/day11.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 11 solution 1")
	fmt.Println(input)
	fmt.Println("day 11 solution 2")
	fmt.Println()
}
