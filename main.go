package main

import (
	"advent2023/day1"
	"advent2023/day2"
	"os"
)

type Solver interface {
	Solve()
}

func main() {
	day := os.Args[1]
	solvers := map[string]Solver{
		"day1": day1.Solver{},
		"day2": day2.Solver{},
	}
	solver, ok := solvers[day]
	if !ok {
		println("solver not found for")
		println(day)
		os.Exit(1)
	}
	solver.Solve()
}
