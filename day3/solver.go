package day3

import (
	"advent2023/mathx"
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	rawScematic, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 3 solution 1")
	fmt.Println(
		mathx.Sum(
			PartNumberFinder{}.Find(
				ParseSchematic(bytes.NewReader(rawScematic)),
			),
		),
	)
	fmt.Println("day 3 solution 2")
}
