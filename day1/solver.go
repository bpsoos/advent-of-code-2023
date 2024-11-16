package day1

import (
	"bytes"
	"os"
)

type Solver struct{}

func (solver Solver) Solve() {
	contents, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	println("day 1 solution 1:")
	println(Decrypter{}.Decrypt(bytes.NewReader(contents)))
	println("day 1 solution 2:")
	println(Decrypter{}.DecryptWithLetters(bytes.NewReader(contents)))
}
