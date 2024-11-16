package day2

import (
	"advent2023/mathx"
	"bytes"
	"fmt"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	contents, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 2 solution 1:")
	fmt.Println(
		mathx.Sum(
			GameVerifier{
				BagContent: CubeSet{
					Red:   12,
					Green: 13,
					Blue:  14,
				},
			}.Verify(
				ParseGames(bytes.NewReader(contents)),
			),
		),
	)
	fmt.Println("day 2 solution 2:")
	fmt.Println(
		calculatePartTwoResults(
			MinSetFinder{}.Find(
				ParseGames(bytes.NewReader(contents)),
			),
		),
	)
}

func calculatePartTwoResults(minSets MinSets) int {
	result := 0
	for _, minSet := range minSets {
		result += minSet.Red * minSet.Green * minSet.Blue
	}

	return result
}
