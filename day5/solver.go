package day5

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

type Solver struct{}

func (Solver) Solve() {
	testInput, err := os.ReadFile("inputs/day5_test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 5 test solution")
	fmt.Println(
		slices.Min(
			LocationMapper{
				Almanac: ParseAlmanac(bytes.NewReader(testInput)),
			}.Locations(),
		),
	)

	input, err := os.ReadFile("inputs/day5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("day 5 solution 1")
	fmt.Println(
		slices.Min(
			LocationMapper{
				Almanac: ParseAlmanac(bytes.NewReader(input)),
			}.Locations(),
		),
	)
	fmt.Println("day 5 solution 2")
}
