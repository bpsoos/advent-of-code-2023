package day4

import (
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testCards, err := os.ReadFile("inputs/day4_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 4 test solution")
	fmt.Println(
		PointCalculator{}.Calculate(
			ParseCards(bytes.NewReader(testCards)),
		),
	)

	scratchcards, err := os.ReadFile("inputs/day4.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 4 solution 1")
	fmt.Println(
		PointCalculator{}.Calculate(
			ParseCards(bytes.NewReader(bytes.TrimSpace(scratchcards))),
		),
	)
}
