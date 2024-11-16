package day2

import (
	"advent2023/math"
	"bytes"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	contents, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	println("day 2 solution 1:")
	possibleGames := GameVerifier{}.Verify(bytes.NewReader(contents))
	println(math.Sum(possibleGames))
}
