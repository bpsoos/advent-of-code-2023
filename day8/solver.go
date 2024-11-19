package day8

import (
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	_, err := os.ReadFile("inputs/day8_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 8 test solution 1")
	fmt.Println()
	fmt.Println("day 8 test solution 2")
	fmt.Println()

	_, err = os.ReadFile("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 8 solution 1")
	fmt.Println()
	fmt.Println("day 8 solution 2")
	fmt.Println()
}
