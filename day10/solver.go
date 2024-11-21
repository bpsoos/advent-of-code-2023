package day10

import (
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	_, err := os.ReadFile("inputs/day10_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 10 test solution 1")
	fmt.Println()
	fmt.Println("day 10 test solution 2")
	fmt.Println()

	_, err = os.ReadFile("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 10 solution 1")
	fmt.Println()
	fmt.Println("day 10 solution 2")
	fmt.Println()
}
