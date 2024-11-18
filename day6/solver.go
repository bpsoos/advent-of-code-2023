package day6

import (
	"advent2023/mathx"
	"fmt"
)

type Solver struct{}

func (Solver) Solve() {
	testInput := []RaceTime{
		{Time: 7, Distance: 9},
		{Time: 15, Distance: 40},
		{Time: 30, Distance: 200},
	}

	fmt.Println("day 6 test solution 1")
	fmt.Println(
		mathx.Multiply(
			NewWinStrategizer(testInput).WinPossibilities(),
		),
	)
	fmt.Println("day 6 test solution 2")
	fmt.Println(
		NewWinStrategizer([]RaceTime{
			{Time: 71530, Distance: 940200},
		}).WinPossibilities()[0],
	)

	input := []RaceTime{
		{Time: 53, Distance: 313},
		{Time: 89, Distance: 1090},
		{Time: 76, Distance: 1214},
		{Time: 98, Distance: 1201},
	}
	fmt.Println("day 6 solution 1")
	fmt.Println(
		mathx.Multiply(
			NewWinStrategizer(input).WinPossibilities(),
		),
	)
	fmt.Println("day 6 solution 2")
	input = []RaceTime{
		{Time: 53897698, Distance: 313109012141201},
	}
	fmt.Println(
		NewWinStrategizer(input).WinPossibilities()[0],
	)
}
