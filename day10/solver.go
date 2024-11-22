package day10

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day9_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 9 test solution 1")
	fmt.Println(Predictor{}.Predict(ParseReport(bytes.NewReader(testInput))))
	fmt.Println("day 9 test solution 2")
	fmt.Println()

	input, err := os.ReadFile("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 9 solution 1")
	fmt.Println(Predictor{}.Predict(ParseReport(bytes.NewReader(input))))
	fmt.Println("day 9 solution 2")
	fmt.Println()
}
