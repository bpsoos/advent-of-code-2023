package main

import (
	"advent2023/day1"
	"advent2023/day10"
	"advent2023/day11"
	"advent2023/day2"
	"advent2023/day3"
	"advent2023/day4"
	"advent2023/day5"
	"advent2023/day6"
	"advent2023/day7"
	"advent2023/day8"
	"advent2023/day9"
	"fmt"
	"os"
)

type Solver interface {
	Solve()
}

func main() {
	day := os.Args[1]
	solvers := map[string]Solver{
		"day1":  day1.Solver{},
		"day2":  day2.Solver{},
		"day3":  day3.Solver{},
		"day4":  day4.Solver{},
		"day5":  day5.Solver{},
		"day6":  day6.Solver{},
		"day7":  day7.Solver{},
		"day8":  day8.Solver{},
		"day9":  day9.Solver{},
		"day10": day10.Solver{},
		"day11": day11.Solver{},
	}
	solver, ok := solvers[day]
	if !ok {
		fmt.Println("solver not found for")
		fmt.Println(day)
		os.Exit(1)
	}
	solver.Solve()
}
