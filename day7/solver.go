package day7

import (
	"advent2023/day7/jokered"
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day7_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 7 test solution 1")
	fmt.Println(
		TotalWinningsCalculator{}.Calculate(ParseGames(bytes.NewReader(testInput))),
	)
	fmt.Println("day 7 test solution 2")
	fmt.Println(
		TotalWinningsCalculator{}.CalculateJokered(jokered.ParseGames(bytes.NewReader(testInput))),
	)

	input, err := os.ReadFile("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 7 solution 1")
	fmt.Println(
		TotalWinningsCalculator{}.Calculate(ParseGames(bytes.NewReader(input))),
	)
	fmt.Println("day 7 solution 2")
	fmt.Println(
		TotalWinningsCalculator{}.CalculateJokered(jokered.ParseGames(bytes.NewReader(input))),
	)
}
