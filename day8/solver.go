package day8

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day8_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 8 test solution 1")
	pattern, treeMap := ParseInstructions(bytes.NewReader(testInput))
	fmt.Println(
		StepCounter{}.Count(
			pattern, treeMap["AAA"],
		),
	)
	fmt.Println("day 8 test solution 2")
	fmt.Println(
		StepCounter{}.CountAsGhost(
			pattern, treeMap,
		),
	)

	input, err := os.ReadFile("inputs/day8.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 8 solution 1")
	pattern, treeMap = ParseInstructions(bytes.NewReader(input))
	fmt.Println(
		StepCounter{}.Count(
			pattern, treeMap["AAA"],
		),
	)
	fmt.Println("day 8 solution 2")
	fmt.Println(
		StepCounter{}.CountAsGhost(
			pattern, treeMap,
		),
	)
}
