package day5

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day5_test.txt")
	if err != nil {
		panic(err)
	}
	_ = bytes.NewReader(testInput)
	fmt.Println("day 5 test solution")

	input, err := os.ReadFile("inputs/day5.txt")
	_ = bytes.NewReader(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("day 5 solution 1")
	fmt.Println("day 5 solution 2")
}
