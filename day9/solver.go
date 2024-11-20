package day9

import (
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	_, err := os.ReadFile("inputs/day9_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 9 test solution 1")
	fmt.Println()
	fmt.Println("day 9 test solution 2")
	fmt.Println()

	_, err = os.ReadFile("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 9 solution 1")
	fmt.Println()
	fmt.Println("day 9 solution 2")
	fmt.Println()
}
