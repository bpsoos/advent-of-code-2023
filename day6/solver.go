package day6

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day6_test.txt")
	if err != nil {
		panic(err)
	}
	_ = bytes.NewReader(testInput)
	fmt.Println("day 6 test solution 1")
	fmt.Println("day 6 test solution 2")

	input, err := os.ReadFile("inputs/day6.txt")
	_ = bytes.NewReader(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("day 6 solution 1")
	fmt.Println("day 6 solution 2")
}
